package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Sqldb *sql.DB
}

const (
	dataSource = "root:123456@tcp(192.168.1.4:3306)/sidss?charset=utf8"
)

func New() *Database {
	log.Println("[INFO]", "Open database...")
	sqldb, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalln("[FATAL]", "model.New()", "sql.Open()", err)
		return nil
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatalln("[FATAL]", err)
		return nil
	}

	log.Println("[INFO]", "Open database...SUCCESS")
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(20)

	return &Database{Sqldb: sqldb}
}

var std = New()

func SqlDB() *sql.DB {
	return std.Sqldb
}
