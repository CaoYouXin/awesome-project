package db

import (
	"database/sql"
	"embed"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed db.sqlite
var db embed.FS

var Inst *sql.DB

func InitInst() error {
	dbConn, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return err
	}

	dbContent, err := db.ReadFile("db.sqlite")
	if err != nil {
		return err
	}

	_, err = dbConn.Exec(string(dbContent))
	if err != nil {
		return err
	}

	Inst = dbConn

	return nil
}
