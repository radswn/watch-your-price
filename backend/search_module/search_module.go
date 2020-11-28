package search_module

import (
	"backend/search_module/website_type"
	"errors"
)

type SearchResult struct {
	Phrase     string
	Page       int
	NumOfPages int
	Results    map[string]string
}

type SearchRequest struct {
	Phrase  string
	Page    int
	Website website_type.WebsiteType
}

type WebsiteSearch interface {
	GetResults(phrase string, page int) (SearchResult, error)
}

type searchModule struct {
	websites map[website_type.WebsiteType]WebsiteSearch
}

func New(websites map[website_type.WebsiteType]WebsiteSearch) (searchModule, error) {
	if len(websites) == 0 {
		return searchModule{}, errors.New("Search module should have at least one website")
	}
	search := searchModule{websites: websites}
	return search, nil
}
