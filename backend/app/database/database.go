package database

import (
	"database/sql"
	"fmt"

	"github.com/igroks/sd-project/backend/app/config"
	_ "github.com/lib/pq"
)

func OpenConn(database config.DatabaseConfig) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Database,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
