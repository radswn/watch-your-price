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
	Phrase  string                   `json:"phrase" binding:"required"`
	Page    int                      `json:"page" binding:"required"`
	Website website_type.WebsiteType `json:"website" binding:"required"`
}

type WebsiteSearch interface {
	GetResults(phrase string, page int) (SearchResult, error)
}

type SearchModule struct {
	websites map[website_type.WebsiteType]WebsiteSearch
}

func New(websites map[website_type.WebsiteType]WebsiteSearch) (*SearchModule, error) {
	if len(websites) == 0 {
		return nil, errors.New("Search module should have at least one website")
	}
	search := &SearchModule{websites: websites}
	return search, nil
}

func (sm SearchModule) Search(request SearchRequest) (*SearchResult, error) {
	result, err := sm.websites[request.Website].GetResults(request.Phrase, request.Page)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
