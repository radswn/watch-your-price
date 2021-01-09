package websites

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const listViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_list</title>\n</head>\n<body>\n<strong class=\"cat-prod-row__name\"><a href=\"/121\">Product 1</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/122\" class=\"go-to-shop\">Product 2</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/123\">Product 3</a></strong>\n</body>\n</html>\n"

const gridViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_grid</title>\n</head>\n<body>\n<div class=\"grid-row\">\n    <a href=\"/121\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 1</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/122\" class=\"go-to-shop\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 2</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/123\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 3</strong>\n    </div>\n</div>\n</body>\n</html>"

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
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldReturnURLWithPageInformationWhenPageIsMoreThanZero() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldHaveGivenPhraseAndPageInResult() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldReturnAnyResultsInListView() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldReturnAnyResultsInGridView() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldOmitAnyProductThatLinksToExternalPageInListView() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldOmitAnyProductThatLinksToExternalPageInGridView() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestLinksInResultShouldLinkToCeneo() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldReadMaxNumOfPages() {
	suite.T().FailNow()
}

func (suite ceneoTestSuite) TestShouldDefaultMaxNumOfPagesToZeroWhenThereIsOnlyOnePage() {
	suite.T().FailNow()
}

func (suite *ceneoTestSuite) TestShouldOmitResultsThatHaveNoLink() {
	suite.T().FailNow()
}
