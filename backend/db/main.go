package main

import (
	"database/sql"

	"github.com/pkl-min-1/discusspedia/backend/db/migration"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "basis-app.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	migration.Migrate(db)
}
