package check

import (
	"errors"
	"github.com/sirupsen/logrus"
	"search_module/scraper"
)

// Result represent the response from the website price check
type Result struct {
	Price string `json:"price" binding:"required"`
}

// Request represent query to the specific website price check
type Request struct {
	Url     string              `json:"phrase" binding:"required"`
	Website scraper.WebsiteType `json:"website" binding:"required"`
}

// WebsiteCheck defines interface that has to be implemented by any website check
type WebsiteCheck interface {
	GetResults(url string) (Result, error)
}

// Module represent struct used to execute methods related to price checking
type Module struct {
	websites map[scraper.WebsiteType]WebsiteCheck
}

// NewCheck returns new instance of Module with provided websites
func NewCheck(websites map[scraper.WebsiteType]WebsiteCheck) (*Module, error) {
	if len(websites) == 0 {
		return nil, errors.New("check module should have at least one website")
	}
	check := &Module{websites: websites}
	return check, nil
}

// Check takes Request parameter, performs query to specific website data and returns results
func (cm Module) Check(request Request) (*Result, error) {
	result, err := cm.websites[request.Website].GetResults(request.Url)
	if err != nil {
		logrus.WithError(err).Error("Could not get results from webpage")
		return nil, err
	}
	return &result, nil
}
