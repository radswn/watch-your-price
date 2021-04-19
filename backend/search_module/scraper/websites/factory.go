package websites

import (
	"search_module/scraper"
)

// NewSearch returns new instance of SearchModule with provided websites
func NewSearch(websiteType scraper.WebsiteType) scraper.WebsiteSearch {
	var module scraper.WebsiteSearch
	switch websiteType {
	case scraper.Ceneo:
		module = newCeneoSearch()
	}
	return module
}

// NewCheck returns new instance of CheckModule with provided websites
func NewCheck(websiteType scraper.WebsiteType) scraper.WebsiteCheck {
	var module scraper.WebsiteCheck
	switch websiteType {
	case scraper.Ceneo:
		module = newCeneoCheck()
	}
	return module
}
