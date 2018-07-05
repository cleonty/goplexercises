package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type ExtractRule struct {
	RelativePath string `json:"path"`
	Attribute    string `json:"attr,omitempty"`
}

type ParseRule struct {
	URL              string      `json:"url"`
	NewsElementsPath string      `json:"newsPath"`
	LinkRule         ExtractRule `json:"linkRule"`
	TitleRule        ExtractRule `json:"titleRule"`
	Interval         uint        `json:"interval"`
}

type Item struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}

type App struct {
	db         *sql.DB
	parseRules []ParseRule
}

func (app *App) readRules() error {
	data, err := ioutil.ReadFile("./rules.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &app.parseRules); err != nil {
		return fmt.Errorf("error while reading parsing rules: %v\n", err)
	}
	return nil
}

func loadHTMLItems(rule *ParseRule) ([]Item, error) {
	var items []Item
	doc, err := htmlquery.LoadURL(rule.URL)
	if err != nil {
		return nil, err
	}
	for _, node := range htmlquery.Find(doc, rule.NewsElementsPath) {
		link := extractEntity(node, &rule.LinkRule)
		title := extractEntity(node, &rule.TitleRule)
		link = processURL(rule.URL, link)
		item := Item{
			Link:  link,
			Title: title,
		}
		items = append(items, item)
	}
	return items, nil
}

func extractEntity(parentNode *html.Node, rule *ExtractRule) string {
	var result string
	node := htmlquery.FindOne(parentNode, rule.RelativePath)
	if node != nil {
		if rule.Attribute != "" {
			result = htmlquery.SelectAttr(node, rule.Attribute)
		} else {
			result = htmlquery.InnerText(node)
		}
	}
	return result
}

func processURL(baseURL string, linkURL string) string {
	url, err := url.Parse(linkURL)
	if err != nil {
		return ""
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	if !url.IsAbs() {
		return base.ResolveReference(url).String()
	}
	return linkURL
}

func readRules() ([]ParseRule, error) {
	data, err := ioutil.ReadFile("./rules.json")
	if err != nil {
		return nil, err
	}
	var parseRules []ParseRule
	if err = json.Unmarshal(data, &parseRules); err != nil {
		return nil, fmt.Errorf("Сбой маршалинга JSON: %v\n", err)
	}
	return parseRules, nil
}

func (app *App) SearchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}
		query := r.Form.Get("q")
		items, err := getNews(app.db, query)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
		}
		data, err := json.MarshalIndent(items, "", "")
		if err != nil {
			log.Fatalf("Сбой маршалинга JSON: %v\n", err)
		}
		w.Header().Set("Content-type", "application/json")
		fmt.Fprintf(w, "%s\n", data)
	})
}

func updateNewsPeriodically(app *App, rule ParseRule) {
	updateNews(app, rule)
	ticker := time.NewTicker(time.Duration(rule.Interval) * time.Minute)
	for {
		select {
		case <-ticker.C:
			updateNews(app, rule)
		}
	}
}

func updateNews(app *App, rule ParseRule) {
	items, err := loadHTMLItems(&rule)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		//fmt.Printf("%d %s(%s)\n", i, item.Title, item.Link)
		err = insertNews(app.db, &item)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	app := &App{}
	err := app.readRules()
	if err != nil {
		log.Fatal(err)
	}
	rules := app.parseRules
	fmt.Printf("%+v\n", rules)
	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	app.db = db
	for _, rule := range rules {
		go updateNewsPeriodically(app, rule)
	}
	fmt.Println("Новости")
	items, err := getNews(db, "")
	if err != nil {
		log.Fatal(err)
	}
	for i, item := range items {
		fmt.Printf("%d %s(%s)\n", i, item.Title, item.Link)
	}

	http.Handle("/news/", app.SearchHandler())
	http.Handle("/", http.FileServer(http.Dir("./client/dist/client")))
	http.ListenAndServe(":8080", nil)
}
