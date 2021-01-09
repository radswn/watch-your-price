package websites

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

const listViewHtmlWithoutPage = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_list</title>\n</head>\n<body>\n<strong class=\"cat-prod-row__name\"><a href=\"/121\">Product 1</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/122\" class=\"go-to-shop\">Product outside</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/123\">Product 3</a></strong>\n</body>\n</html>\n"

const listViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_list</title>\n</head>\n<body>\n<strong class=\"cat-prod-row__name\"><a href=\"/121\">Product 1</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/122\" class=\"go-to-shop\">Product 2</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/123\">Product 3</a></strong>\n<strong class=\"cat-prod-row__name\">ProductWithoutLink</strong>\n<div class=\"pagination-top\">\n    <input id=\"page-counter\" data-pageCount=\"824\">\n    z 824\n</div>\n</body>\n</html>"

const gridViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_grid</title>\n</head>\n<body>\n<div class=\"grid-row\">\n    <a href=\"/121\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 1</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/122\" class=\"go-to-shop\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product outside</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/123\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 3</strong>\n    </div>\n</div>\n</body>\n</html>"

type ceneoTestSuite struct {
	suite.Suite
	server      *httptest.Server
	ceneoSearch *ceneoSearch
}

func (suite *ceneoTestSuite) SetupSuite() {
	suite.server = testServer()
	suite.ceneoSearch = testCeneoSearch(suite.server.URL)
}

func testServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(listViewHtml))
	})

	mux.HandleFunc("/listWithoutPages", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(listViewHtmlWithoutPage))
	})

	mux.HandleFunc("/grid", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(gridViewHtml))
	})
	return httptest.NewServer(mux)
}

func testCeneoSearch(serverURL string) *ceneoSearch {
	domain := strings.TrimPrefix(serverURL, "http://")
	domain = removePort(domain)
	return &ceneoSearch{
		queueStorage: 10,
		queueThreads: 1,
		domain:       domain,
		domainGlob:   "*",
		baseUrl:      serverURL,
		delay:        0,
		randomDelay:  0,
	}
}

func removePort(domain string) string {
	return strings.Split(domain, ":")[0]
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(ceneoTestSuite))
}

func (suite *ceneoTestSuite) TestShouldReturnURLWithoutPageWhenPageIsZero() {
	phrase := "Product"
	page := 0
	baseUrl := "www.abc.com/;search-"
	ceneoSearch := &ceneoSearch{baseUrl: baseUrl}
	url := ceneoSearch.createSearchUrl(phrase, page)

	assert.Equal(suite.T(), "www.abc.com/;search-"+phrase+".htm?nocatnarrow=1", url)
}

func (suite *ceneoTestSuite) TestShouldReturnURLWithPageInformationWhenPageIsMoreThanZero() {
	phrase := "Product"
	page := 3
	baseUrl := "www.abc.com/;search-"
	ceneoSearch := &ceneoSearch{baseUrl: baseUrl}
	url := ceneoSearch.createSearchUrl(phrase, page)

	assert.Equal(suite.T(), "www.abc.com/;search-"+phrase+";0020-30-0-0-"+strconv.Itoa(page)+".htm?nocatnarrow=1", url)
}

func (suite *ceneoTestSuite) TestShouldHaveGivenPhraseAndPageInResult() {
	phrase := "Product"
	page := 3
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", phrase, page)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result.Phrase, phrase)
	assert.Equal(suite.T(), result.Page, page)
}

func (suite *ceneoTestSuite) TestShouldReturnAnyResultsInListView() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", "", 0)
	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(result.Results), 0)
}

func (suite *ceneoTestSuite) TestShouldReturnAnyResultsInGridView() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/grid", "", 0)
	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(result.Results), 0)
}

func (suite *ceneoTestSuite) TestShouldOmitAnyProductThatLinksToExternalPageInListView() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", "", 0)
	assert.Nil(suite.T(), err)
	assert.NotContains(suite.T(), result.Results, "Product outside")
}

func (suite *ceneoTestSuite) TestShouldOmitAnyProductThatLinksToExternalPageInGridView() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/grid", "", 0)
	assert.Nil(suite.T(), err)
	assert.NotContains(suite.T(), result.Results, "Product outside")
}

func (suite *ceneoTestSuite) TestLinksInResultShouldLinkToTheSameDomain() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", "", 0)
	assert.Nil(suite.T(), err)
	for _, link := range result.Results {
		assert.True(suite.T(), strings.HasPrefix(link, "http://"+suite.ceneoSearch.domain))
	}
}

func (suite *ceneoTestSuite) TestShouldReadMaxNumOfPages() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", "", 0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 824, result.NumOfPages)
}

func (suite ceneoTestSuite) TestShouldDefaultMaxNumOfPagesToZeroWhenThereIsOnlyOnePage() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/listWithoutPages", "", 0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, result.NumOfPages)
}

func (suite *ceneoTestSuite) TestShouldOmitResultsThatHaveNoLink() {
	result, err := suite.ceneoSearch.search(suite.server.URL+"/list", "", 0)
	assert.Nil(suite.T(), err)
	assert.NotContains(suite.T(), result.Results, "ProductWithoutLink")
}
