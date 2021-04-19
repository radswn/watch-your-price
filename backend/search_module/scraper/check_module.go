package scraper

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// CheckResult represent the response from the website price check
type CheckResult struct {
	Price string `json:"price" binding:"required"`
}

// CheckRequest represent query to the specific website price check
type CheckRequest struct {
	Url     string      `json:"phrase" binding:"required"`
	Website WebsiteType `json:"website" binding:"required"`
}

// WebsiteCheck defines interface that has to be implemented by any website check
type WebsiteCheck interface {
	GetResults(url string) (CheckResult, error)
}

// CheckModule represent struct used to execute methods related to price checking
type CheckModule struct {
	websites map[WebsiteType]WebsiteCheck
}

// NewCheck returns new instance of CheckModule with provided websites
func NewCheck(websites map[WebsiteType]WebsiteCheck) (*CheckModule, error) {
	if len(websites) == 0 {
		return nil, errors.New("check module should have at least one website")
	}
	check := &CheckModule{websites: websites}
	return check, nil
}

// Check takes CheckRequest parameter, performs query to specific website data and returns results
func (cm CheckModule) Check(request CheckRequest) (*CheckResult, error) {
	result, err := cm.websites[request.Website].GetResults(request.Url)
	if err != nil {
		logrus.WithError(err).Error("Could not get results from webpage")
		return nil, err
	}
	return &result, nil
}
