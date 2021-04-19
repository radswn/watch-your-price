package websites

import (
	"search_module/search"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType search.WebsiteType) search.WebsiteSearch {
	var module search.WebsiteSearch
	switch websiteType {
	case search.Ceneo:
		module = newCeneoSearch()
	}
	return module
}
