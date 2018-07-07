package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/antchfx/htmlquery"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/html"
)

const databseFile = "./news.db"
const parsingRulesFile = "./rules.json"

type ExtractRule struct {
	RelativePath string `json:"path"`
	Attribute    string `json:"attr,omitempty"`
}

type ParsingRule struct {
	Interval         uint        `json:"interval"`
	URL              string      `json:"url"`
	NewsElementsPath string      `json:"newsPath"`
	LinkRule         ExtractRule `json:"linkRule"`
	TitleRule        ExtractRule `json:"titleRule"`
}

// NewsItem represnts a news
type NewsItem struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}

type App struct {
	db           *sql.DB
	server       *http.Server
	parsingRules []ParsingRule
	wg           sync.WaitGroup
}

func (app *App) readParsingRules() error {
	data, err := ioutil.ReadFile(parsingRulesFile)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &app.parsingRules); err != nil {
		return fmt.Errorf("error while reading parsing rules: %v", err)
	}
	return nil
}

func (app *App) loadNewsList(rule *ParsingRule) ([]NewsItem, error) {
	var items []NewsItem
	doc, err := htmlquery.LoadURL(rule.URL)
	if err != nil {
		return nil, err
	}
	for _, node := range htmlquery.Find(doc, rule.NewsElementsPath) {
		link := extractEntity(node, &rule.LinkRule)
		title := extractEntity(node, &rule.TitleRule)
		link, err = convertToAbsURL(rule.URL, link)
		if err != nil {
			return nil, fmt.Errorf("error converting link url %s to absolute url using base url %s: %v", link, rule.URL, err) 
		}
		item := NewsItem{
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

func convertToAbsURL(baseURL string, linkURL string) (string, error) {
	url, err := url.Parse(linkURL)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	if !url.IsAbs() {
		return base.ResolveReference(url).String(), nil
	}
	return linkURL, nil
}

func (app *App) searchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}
		query := r.Form.Get("q")
		items, err := app.getNews(query)
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
		}
		data, err := json.MarshalIndent(items, "", "")
		if err != nil {
			return
		}
		w.Header().Set("Content-type", "application/json")
		fmt.Fprintf(w, "%s\n", data)
	})
}

func (app *App) updateNewsPeriodically(ctx context.Context, rule ParsingRule) {
	app.updateNews(rule)
	defer app.wg.Done()
	ticker := time.NewTicker(time.Duration(rule.Interval) * time.Minute)
	for {
		select {
		case <-ticker.C:
			app.updateNews(rule)
		case <-ctx.Done():
			log.Printf("Stop updater for %s\n", rule.URL)
			return
		}
	}
}

func (app *App) updateNews(rule ParsingRule) {
	items, err := app.loadNewsList(&rule)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		err = app.insertNewsItem(&item)
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *App) startUpdaters(ctx context.Context) {
	for _, rule := range app.parsingRules {
		app.wg.Add(1)
		go app.updateNewsPeriodically(ctx, rule)
	}
}

func (app *App) start() (chan struct{}, error) {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	app.startUpdaters(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	server := &http.Server{Addr: ":8080"}
	go func() {
		<-c
		log.Println("Exiting...")
		server.Shutdown(ctx)
		cancel()
		app.wg.Wait()
		app.db.Close()
		log.Println("Goodbye!")
		done <- struct{}{}
	}()
	http.Handle("/news/", app.searchHandler())
	http.Handle("/", http.FileServer(http.Dir("./client/dist/client")))
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	return done, nil
}

const newsStatement = `
        CREATE TABLE IF NOT EXISTS 'news' (
        'id' INTEGER PRIMARY KEY AUTOINCREMENT,
        'link' VARCHAR(1024) UNIQUE NOT NULL,
        'title' VARCHAR(1024) NOT NULL,
		'timestamp' DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

func (app *App) openDatabase() error {
	db, err := sql.Open("sqlite3", databseFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(newsStatement)
	if err != nil {
		db.Close()
		return err
	}
	app.db = db
	return nil
}

func (app *App) getNews(query string) ([]NewsItem, error) {
	items := make([]NewsItem, 0)
	var statement string
	if query != "" {
		statement = "SELECT link, title FROM news WHERE instr(title, ?) <> 0 ORDER BY timestamp DESC"
	} else {
		statement = "SELECT link, title FROM news ORDER BY timestamp DESC"
	}
	rows, err := app.db.Query(statement, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item NewsItem
		err = rows.Scan(&item.Link, &item.Title)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (app *App) insertNewsItem(item *NewsItem) error {
	_, err := app.db.Exec("INSERT INTO news(link, title) values(?, ?)", item.Link, item.Title)
	if err != nil {
		return fmt.Errorf("Insert failed for link='%s', title='%s': %v", item.Link, item.Title, err)
	}
	return nil
}

func main() {
	app := &App{}
	err := app.readParsingRules()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", app.parsingRules)
	err = app.openDatabase()
	if err = app.openDatabase(); err != nil {
		log.Fatal(err)
	}
	done, err := app.start()
	if err != nil {
		log.Fatal(err)
	}
	<-done
	fmt.Println("Done!")
}
