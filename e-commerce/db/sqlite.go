package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Con *sql.DB
}

func OpenSqlite() Sqlite {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(1)

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return Sqlite{Con: db}
}

func (s *Sqlite) GetOne() error {
	log.Print("GetOne")
	return nil
}
