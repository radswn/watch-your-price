package check

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const itemHtml = "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <title>Xiaomi Mi 10T Pro 5G 8/256GB Srebrny - Cena, opinie na Ceneo.pl</title>\n    <meta property=\"product:price:currency\" content=\"PLN\" />\n    <meta property=\"product:price:amount\" content=\"2246.58\" />\n    <meta property=\"og:url\" content=\"https://www.ceneo.pl/98016017\" />\n</head>\n<body>\n</body>\n</html>"

type ceneoTestSuite struct {
	suite.Suite
	server     *httptest.Server
	ceneoCheck *ceneoCheck
}

func (suite *ceneoTestSuite) SetupSuite() {
	suite.server = testCheckServer()
	suite.ceneoCheck = testCeneoCheck(suite.server.URL)
}

func testCheckServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(itemHtml))
	})

	return httptest.NewServer(mux)
}

func testCeneoCheck(serverURL string) *ceneoCheck {
	domain := strings.TrimPrefix(serverURL, "http://")
	domain = removePort(domain)
	return &ceneoCheck{
		queueStorage: 10,
		queueThreads: 1,
		domain:       domain,
		domainGlob:   "*",
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
