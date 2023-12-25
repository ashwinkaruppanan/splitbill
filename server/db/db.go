package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("postgres", os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ping to database...")

	return &DB{
		db,
	}
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func (d *DB) CloseDB() {
	d.db.Close()
}
