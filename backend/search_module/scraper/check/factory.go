package check

import "search_module/scraper"

// NewWebsiteCheck returns new instance of WebsiteCheck
func NewWebsiteCheck(websiteType scraper.WebsiteType) WebsiteCheck {
	var module WebsiteCheck
	switch websiteType {
	case scraper.Ceneo:
		module = newCeneoCheck()
	}
	return module
}
