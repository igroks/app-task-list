package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn(url string) *sql.DB {

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
