package websites

import (
	"search_module"
	"search_module/website_type"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType website_type.WebsiteType) search_module.WebsiteSearch {
	var module search_module.WebsiteSearch
	switch websiteType {
	case website_type.Ceneo:
		module = newCeneoSearch()
	}
	return module
}
