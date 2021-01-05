package websites

import (
	"backend/search_module"
	"backend/search_module/website_type"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType website_type.WebsiteType) search_module.WebsiteSearch {
	var module search_module.WebsiteSearch
	switch websiteType {
	case website_type.Ceneo:
		module = &ceneoSearch{}
	}
	return module
}
