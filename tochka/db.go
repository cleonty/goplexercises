package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const newsStatement = `
        CREATE TABLE IF NOT EXISTS news (
        'id' INTEGER PRIMARY KEY AUTOINCREMENT,
        'link' VARCHAR(1024) UNIQUE NOT NULL,
        'title' VARCHAR(1024) NOT NULL,
		'timestamp' DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

func createDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./news.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(newsStatement)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func getNews(db *sql.DB, query string) ([]Item, error) {
	var items []Item
	rows, err := db.Query("SELECT link, title FROM news WHERE instr(title, ?) <> 0 ORDER BY timestamp DESC", query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item Item
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

func insertNews(db *sql.DB, item *Item) error {
	_, err := db.Exec("INSERT INTO news(link, title) values(?, ?)", item.Link, item.Title)
	if err != nil {
		return fmt.Errorf("Insert failed link='%s', title='%s': %v", item.Link, item.Title, err)
	}
	return nil
}
