package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"search_module/scraper"
)
import _ "github.com/mattn/go-sqlite3"

type Checker struct {
	database      *sql.DB
	scraperModule *scraper.Module
}

type Product struct {
	Id    string
	Link  string
	Price string
}

func NewDatabaseChecker(module *scraper.Module) *Checker {
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
	return &Checker{database: db, scraperModule: module}
}

func (c *Checker) GetAllProducts() []Product {
	var result []Product

	rows, err := c.database.Query("SELECT id, link, price FROM entities_product")
	if err != nil {
		logrus.WithError(err).Warn("Cannot get products from database")
		return result
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Link, &product.Price)
		if err != nil {
			logrus.WithError(err).Warn("Cannot get information for some products")
		}
		result = append(result, product)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithError(err).Warn("Cannot get information for some products")
		return []Product{}
	}

	return result
}

func (c *Checker) CloseDatabase() {
	err := c.database.Close()
	if err != nil {
		logrus.WithError(err).Warn("Cannot close database")
	}
}
