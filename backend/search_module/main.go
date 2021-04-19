package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
	"search_module/scraper"
	"search_module/scraper/websites"
	"strconv"
	"strings"
)

var searchModule *scraper.SearchModule
var checkModule *scraper.CheckModule

func init() {
	setupLogrus()
	searchModule = setupSearchModule()
	checkModule = setupCheckModule()

	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")
	r.Handle("/check", http.HandlerFunc(checkHandler)).Methods("GET")

	logrus.Fatal(http.ListenAndServe(":8001", r))
}

func main() {

}

func setupLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(r *runtime.Frame) (function string, file string) {
			filepath := strings.Split(r.File, "/")
			return "", fmt.Sprintf("%s:%v", filepath[len(filepath)-1], r.Line)
		},
	})

	file, err := os.OpenFile("search.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.SetOutput(os.Stdout)
		logrus.WithError(err).Warn("Cannot open log file. Logging to stdout.")
	} else {
		logrus.SetOutput(file)
	}

	// adds information about location of log
	logrus.SetReportCaller(true)
}

func setupSearchModule() *scraper.SearchModule {
	ceneoSearch := websites.NewSearch(scraper.Ceneo)
	searchModule, err := scraper.NewSearch(map[scraper.WebsiteType]scraper.WebsiteSearch{
		scraper.Ceneo: ceneoSearch,
	})
	if err != nil {
		logrus.WithError(err).Panic("Can't initialize search module.")
	}
	return searchModule
}

func setupCheckModule() *scraper.CheckModule {
	ceneoCheck := websites.NewCheck(scraper.Ceneo)
	checkModule, err := scraper.NewCheck(map[scraper.WebsiteType]scraper.WebsiteCheck{
		scraper.Ceneo: ceneoCheck,
	})
	if err != nil {
		logrus.WithError(err).Panic("Can't initialize search module.")
	}
	return checkModule
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	queryParameters := r.URL.Query()
	phraseQuery, ok := queryParameters["phrase"]
	if !ok || len(phraseQuery) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	phrase := phraseQuery[0]

	websiteQuery, ok := queryParameters["website"]
	if !ok || len(websiteQuery) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	website, err := convertWebsite(websiteQuery[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var page int
	pageQuery, ok := queryParameters["page"]
	if ok && len(pageQuery) >= 1 {
		page, _ = strconv.Atoi(pageQuery[0])
	}

	request := scraper.SearchRequest{Page: page, Phrase: phrase, Website: website}
	result, _ := searchModule.Search(request)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	queryParameters := r.URL.Query()
	urlQuery, ok := queryParameters["url"]
	if !ok || len(urlQuery) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url := urlQuery[0]

	websiteQuery, ok := queryParameters["website"]
	if !ok || len(websiteQuery) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	website, err := convertWebsite(websiteQuery[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := scraper.CheckRequest{Url: url, Website: website}
	result, _ := checkModule.Check(request)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func convertWebsite(websiteStr string) (scraper.WebsiteType, error) {
	var website scraper.WebsiteType
	switch strings.ToLower(websiteStr) {
	case "ceneo":
		website = scraper.Ceneo
		break
	default:
		return "", errors.New("unknown website")
	}
	return website, nil
}
