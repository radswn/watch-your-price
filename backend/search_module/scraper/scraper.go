package scraper

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// WebsiteScraper defines interface that has to be implemented by any website module
type WebsiteScraper interface {
	Search(phrase string, page int) (SearchResult, error)
	CheckPrice(url string) (CheckResult, error)
}

// Module represent struct used to execute methods related to searching
type Module struct {
	websites map[WebsiteType]WebsiteScraper
}

// New returns new instance of Module with provided websites
func New(websites map[WebsiteType]WebsiteScraper) (*Module, error) {
	if len(websites) == 0 {
		return nil, errors.New("search module should have at least one website")
	}
	search := &Module{websites: websites}
	return search, nil
}

// Search takes SearchRequest parameter, performs query to specific website data and returns results
func (m Module) Search(request SearchRequest) (*SearchResult, error) {
	result, err := m.websites[request.Website].Search(request.Phrase, request.Page)
	if err != nil {
		logrus.WithError(err).Error("Could not perform search on webpage")
		return nil, err
	}
	return &result, nil
}

// CheckPrice takes Request parameter, performs query to specific website data and returns results
func (m Module) CheckPrice(request CheckRequest) (*CheckResult, error) {
	result, err := m.websites[request.Website].CheckPrice(request.Url)
	if err != nil {
		logrus.WithError(err).Error("Could not check price on webpage")
		return nil, err
	}
	return &result, nil
}
