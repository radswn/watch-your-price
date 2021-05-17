package scraper_test

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"search_module/scraper"
	"testing"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

type testWebsiteScraper struct {
}

func (tws testWebsiteScraper) Search(phrase string, page int) (scraper.SearchResult, error) {
	sr := scraper.SearchResult{
		Phrase:     phrase,
		Page:       page,
		NumOfPages: 5,
		Results: map[string]string{
			"result1": "example.com/1",
			"result2": "example.com/2",
			"result3": "example.com/3",
			"result4": "example.com/4",
		},
	}
	return sr, nil
}

func (tws testWebsiteScraper) CheckPrice(url string) (scraper.CheckResult, error) {
	result := scraper.CheckResult{
		Price: "10",
	}
	return result, nil
}

func TestSearchShouldReturnResultsFromWebsiteSearchImplementation(t *testing.T) {
	websiteScraperMap := make(map[scraper.WebsiteType]scraper.WebsiteScraper)
	websiteScraperMap[scraper.Ceneo] = testWebsiteScraper{}
	module, err := scraper.New(websiteScraperMap)
	assert.Nil(t, err)

	requestData := scraper.SearchRequest{
		Phrase:  "test",
		Page:    3,
		Website: "ceneo",
	}

	expectedPhrase := "test"
	expectedPage := 3
	expectedResults := map[string]string{
		"result1": "example.com/1",
		"result2": "example.com/2",
		"result3": "example.com/3",
		"result4": "example.com/4",
	}

	result, err := module.Search(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPhrase, result.Phrase)
	assert.Equal(t, expectedPage, result.Page)
	assert.Equal(t, expectedResults, result.Results)
}

func TestCheckShouldReturnResultsFromWebsiteCheckImplementation(t *testing.T) {
	websiteScraperMap := make(map[scraper.WebsiteType]scraper.WebsiteScraper)
	websiteScraperMap[scraper.Ceneo] = testWebsiteScraper{}
	module, err := scraper.New(websiteScraperMap)
	assert.Nil(t, err)

	requestData := scraper.CheckRequest{
		Url:     "example.com/3",
		Website: "ceneo",
	}

	expectedPrice := "10"

	result, err := module.CheckPrice(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPrice, result.Price)
}
