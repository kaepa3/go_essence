package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors(aushor_id TEXT, author TEXT, PRIMARY KEY (author_id))`,
		`CREATE TABLE IF NOT EXISTS contents(aushor_id TEXT, title_id TEXT, title TEXT, content TEXT, PRIMARY KEY (author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}

	}
}
