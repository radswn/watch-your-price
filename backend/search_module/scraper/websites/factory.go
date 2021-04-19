package websites

import (
	"search_module/scraper"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType scraper.WebsiteType) scraper.WebsiteSearch {
	var module scraper.WebsiteSearch
	switch websiteType {
	case scraper.Ceneo:
		module = newCeneoSearch()
	}
	return module
}
