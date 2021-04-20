package scraper_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"search_module/scraper"
	"strings"
	"testing"
)

type ceneoTestSuite struct {
	suite.Suite
	server       *httptest.Server
	ceneoScraper *scraper.CeneoScraper
	itemHtml     string
}

func (suite *ceneoTestSuite) SetupSuite() {
	suite.server = testCeneoServer(suite)
	suite.ceneoScraper = testCeneoScraper(suite.server.URL)
	suite.itemHtml = "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <title>Xiaomi Mi 10T Pro 5G 8/256GB Srebrny - Cena, opinie na Ceneo.pl</title>\n    <meta property=\"product:price:currency\" content=\"PLN\" />\n    <meta property=\"product:price:amount\" content=\"2246.58\" />\n    <meta property=\"og:url\" content=\"https://www.ceneo.pl/98016017\" />\n</head>\n<body>\n</body>\n</html>"
}

func testCeneoServer(suite *ceneoTestSuite) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(suite.itemHtml))
	})

	return httptest.NewServer(mux)
}

func testCeneoScraper(serverURL string) *scraper.CeneoScraper {
	domain := strings.TrimPrefix(serverURL, "http://")
	domain = removePort(domain)
	return &scraper.CeneoScraper{
		Domain:      domain,
		DomainGlob:  "*",
		Delay:       0,
		RandomDelay: 0,
	}
}

func removePort(domain string) string {
	return strings.Split(domain, ":")[0]
}

func TestRunTestCeneoSuite(t *testing.T) {
	suite.Run(t, new(ceneoTestSuite))
}

func (suite *ceneoTestSuite) TestShouldReturnItemPrice() {
	result, err := suite.ceneoScraper.CheckPrice(suite.server.URL + "/item")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "2246.58", result.Price)
}
