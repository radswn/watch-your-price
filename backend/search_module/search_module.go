package search_module

import "errors"

type SearchResult struct {
   Phrase     string
   Page       int
   NumOfPages int
   Results    map[string]string
}
