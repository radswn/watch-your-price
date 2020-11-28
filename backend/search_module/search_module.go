package search_module

import "errors"

type SearchResult struct {
   Phrase     string
   Page       int
   NumOfPages int
   Results    map[string]string
}

type WebsiteSearch interface {
   GetResults(phrase string, page int) (SearchResult, error)
}
