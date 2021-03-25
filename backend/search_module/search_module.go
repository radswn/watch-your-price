package search_module

import (
	"backend/search_module/website_type"
	"errors"

	"github.com/sirupsen/logrus"
)

// SearchResult represent the response from the website search
type SearchResult struct {
	Phrase     string            `json:"phrase" binding:"required"`
	Page       int               `json:"page" binding:"required"`
	NumOfPages int               `json:"numOfPages" binding:"required"`
	Results    map[string]string `json:"results" binding:"required"`
}

// SearchRequest represent query to the specific website search
type SearchRequest struct {
	Phrase  string                   `json:"phrase" binding:"required"`
	Page    int                      `json:"page" binding:"required"`
	Website website_type.WebsiteType `json:"website" binding:"required"`
}

// WebsiteSearch defines interface that has to be implemented by any website search
type WebsiteSearch interface {
	GetResults(phrase string, page int) (SearchResult, error)
}

// SearchModule represent struct used to execute methods related to searching
type SearchModule struct {
	websites map[website_type.WebsiteType]WebsiteSearch
}

// New returns new instance of SearchModule with provided websites
func New(websites map[website_type.WebsiteType]WebsiteSearch) (*SearchModule, error) {
	if len(websites) == 0 {
		return nil, errors.New("search module should have at least one website")
	}
	search := &SearchModule{websites: websites}
	return search, nil
}

// Search takes SearchRequest parameter, performs query to specific website data and returns results
func (sm SearchModule) Search(request SearchRequest) (*SearchResult, error) {
	result, err := sm.websites[request.Website].GetResults(request.Phrase, request.Page)
	if err != nil {
		logrus.WithError(err).Error("Could not get results from webpage")
		return nil, err
	}
	return &result, nil
}
