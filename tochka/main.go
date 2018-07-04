package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"sync"

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
}

type Item struct {
	Link  string
	Title string
}

var rule = ParseRule{
	NewsElementsPath: `//div[@class="item"]`,
	LinkRule: ExtractRule{
		RelativePath: "a",
		Attribute:    "href",
	},
	TitleRule: ExtractRule{
		RelativePath: "a",
		Attribute:    "",
	},
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

func main() {
	rules, err := readRules()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rules)
	var wg sync.WaitGroup
	for _, rule := range rules {
		wg.Add(1)
		go func() {
			defer wg.Done()
			items, err := loadHTMLItems(&rule)
			if err != nil {
				log.Fatal(err)
			}
			for i, item := range items {
				fmt.Printf("%d %s(%s)\n", i, item.Title, item.Link)
			}
		}()
	}
	wg.Wait()

}
