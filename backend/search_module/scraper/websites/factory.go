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
