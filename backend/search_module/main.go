package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
	"search_module/search"
	"search_module/search/website_type"
	"search_module/search/websites"
	"strconv"
	"strings"
)

var searchModule *search.Module

func init() {
	setupLogrus()
	searchModule = setupSearchModule()

	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/", http.HandlerFunc(searchHandler)).Methods("GET")

	logrus.Fatal(http.ListenAndServe(":8000", r))
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

func setupSearchModule() *search.Module {
	ceneoSearch := websites.New(website_type.Ceneo)
	searchModule, err := search.New(map[website_type.WebsiteType]search.WebsiteSearch{
		website_type.Ceneo: ceneoSearch,
	})
	if err != nil {
		logrus.WithError(err).Panic("Can't initialize search module.")
	}
	return searchModule
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

	request := search.Request{Page: page, Phrase: phrase, Website: website}
	result, _ := searchModule.Search(request)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func convertWebsite(s string) (website_type.WebsiteType, error) {

	return website_type.Ceneo, nil
}
