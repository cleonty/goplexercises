package main

import (
	"database/sql"
	"fmt"
	_ "time"

	_ "github.com/mattn/go-sqlite3"
)

var newsStatement = `CREATE TABLE IF NOT EXISTS news (
        'id' INTEGER PRIMARY KEY AUTOINCREMENT,
        'link' VARCHAR(1024) UNIQUE NOT NULL,
        'title' VARCHAR(1024) NOT NULL
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
	rows, err := db.Query("SELECT link, title FROM news WHERE instr(title, ?) <> 0", query)
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
	return items, nil
}

func insertNews(db *sql.DB, item *Item) error {
	stmt, err := db.Prepare("INSERT INTO news(link, title) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(item.Link, item.Title)
	if err != nil {
		return fmt.Errorf("Insert failed link='%s', title='%s': %v", item.Link, item.Title, err)
	}
	return nil
}
