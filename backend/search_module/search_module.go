package search_module

import "errors"

type SearchResult struct {
   Phrase     string
   Page       int
   NumOfPages int
   Results    map[string]string
   Phrase     string
   Page       int
   NumOfPages int
   Results    map[string]string
}

type websiteType string

const (
	ceneo websiteType = "ceneo"
)

type searchRequest struct {
   phrase  string
   page    int
   website websiteType
}

type WebsiteSearch interface {
   GetResults(phrase string, page int) (SearchResult, error)
}

type searchModule struct {
   websites []WebsiteSearch
}

func New(websites []WebsiteSearch) (searchModule, error) {
   if len(websites) == 0 {
       return searchModule{}, errors.New("Search module should have at least one website")
   }
   search := searchModule{websites: websites}
   return search, nil
}
