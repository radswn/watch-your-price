package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)
import _ "github.com/mattn/go-sqlite3"

type Checker struct {
	Database *sql.DB
}

type Product struct {
	Id    string
	Link  string
	Price float32
}

func NewDatabaseChecker() *Checker {
	db, err := sql.Open("sqlite3", "file:../db.sqlite3")
	if err != nil {
		logrus.WithError(err).Fatal("Cannot connect to the database")
		return nil
	}
	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Fatal("Cannot connect to the database")
		return nil
	}
	return &Checker{Database: db}
}
