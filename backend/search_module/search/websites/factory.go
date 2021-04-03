package websites

import (
	"search_module/search"
	"search_module/search/website_type"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType website_type.WebsiteType) search.WebsiteSearch {
	var module search.WebsiteSearch
	switch websiteType {
	case website_type.Ceneo:
		module = newCeneoSearch()
	}
	return module
}
