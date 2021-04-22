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

func (c *Checker) UpdatePrices() {

	products := c.GetAllProducts()
	if len(products) == 0 {
		return
	}

	stmt, err := c.database.Prepare("UPDATE entities_product SET price = ? WHERE id = ?")
	if err != nil {
		logrus.WithError(err).Error("Cannot prepare sql statement")
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	for _, product := range products {
		scraperResult, err := c.scraperModule.CheckPrice(scraper.CheckRequest{
			Url:     product.Link,
			Website: scraper.Ceneo, // TODO change website after adding this field to product
		})
		if err != nil {
			logrus.WithError(err).Warn("Cannot check price for item with url: ", product.Link)
			continue
		}

		if product.Price == scraperResult.Price {
			continue
		}

		_, err = stmt.Exec(scraperResult.Price, product.Id)
		if err != nil {
			logrus.WithError(err).Warn("Cannot update price for item: ", product.Id)
			continue
		}

		logrus.Debugf("Id: %s, old price: %s, new price: %s", product.Id, product.Price, scraperResult.Price)
	}
}

func (c *Checker) CloseDatabase() {
	err := c.database.Close()
	if err != nil {
		logrus.WithError(err).Warn("Cannot close database")
	}
}
