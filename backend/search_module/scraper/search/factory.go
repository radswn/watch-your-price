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

// NewCheck returns new instance of CheckModule with provided websites
//func NewCheck(websiteType scraper.WebsiteType) scraper.WebsiteCheck {
//	var module scraper.WebsiteCheck
//	switch websiteType {
//	case scraper.Ceneo:
//		module = websites.newCeneoCheck()
//	}
//	return module
//}
