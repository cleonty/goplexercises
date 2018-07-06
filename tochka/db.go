package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const databseFile = "./news.db"

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
	var items []NewsItem
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
