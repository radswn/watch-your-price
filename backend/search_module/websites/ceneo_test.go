package websites

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

const listViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_list</title>\n</head>\n<body>\n<strong class=\"cat-prod-row__name\"><a href=\"/121\">Product 1</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/122\" class=\"go-to-shop\">Product 2</a></strong>\n<strong class=\"cat-prod-row__name\"><a href=\"/123\">Product 3</a></strong>\n</body>\n</html>\n"

const gridViewHtml = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>test_grid</title>\n</head>\n<body>\n<div class=\"grid-row\">\n    <a href=\"/121\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 1</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/122\" class=\"go-to-shop\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 2</strong>\n    </div>\n</div>\n<div class=\"grid-row\">\n    <a href=\"/123\"></a>\n    <div class=\"grid-item__caption\">\n        <strong class=\"grid_item__name\">Product 3</strong>\n    </div>\n</div>\n</body>\n</html>"

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

func testCeneoSearch(server *httptest.Server) *ceneoSearch {
	domain := strings.TrimPrefix(server.URL, "http://")
	domain = removePort(domain)
	return &ceneoSearch{
		queueStorage: 10,
		queueThreads: 1,
		domain:       domain,
		domainGlob:   "*",
		baseUrl:      server.URL,
		delay:        0,
		randomDelay:  0,
	}
}

func removePort(domain string) string {
	return strings.Split(domain, ":")[0]
}
