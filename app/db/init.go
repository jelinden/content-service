package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDB(dbFileName string) {
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	createUserTable()
}

func DB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}

func createUserTable() {
	sqlStmt := `
		create table if not exists user (id integer primary key, email text not null unique, password text, apitoken text);
		create table if not exists space (id integer primary key, user_id integer not null, name text not null);
		create table if not exists content (id integer primary key, space_id integer not null, key text not null, value text);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func removeDB(dbFileName string) {
	e := os.Remove(dbFileName)
	if e != nil {
		log.Fatal(e)
	}
}
