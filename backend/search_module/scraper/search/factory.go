package search

import (
	"search_module/scraper"
)

// NewWebsiteSearch returns new instance of WebsiteSearch
func NewWebsiteSearch(websiteType scraper.WebsiteType) WebsiteSearch {
	var module WebsiteSearch
	switch websiteType {
	case scraper.Ceneo:
		module = newCeneoSearch()
	}
	return module
}
