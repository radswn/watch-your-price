package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
	"search_module/database"
	"search_module/scraper"
	"strconv"
	"strings"
)

type AppConfig struct {
	Profile string
}

var scraperModule *scraper.Module
var Config *AppConfig

func init() {
	Config = setupConfig()

	setupLogrus()
	scraperModule = setupScraperModule()

	r := mux.NewRouter().StrictSlash(true)
	r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")
	r.Handle("/check", http.HandlerFunc(checkHandler)).Methods("GET")
	r.Handle("/checkdatabase", http.HandlerFunc(checkDatabaseHandler)).Methods("GET")

	db := database.NewDatabaseChecker()

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			logrus.WithError(err).Warn("Cannot close database")
		}
	}(db.Database)

	logrus.Fatal(http.ListenAndServe(":8001", r))
}

func main() {

}

func setupConfig() *AppConfig {
	var config AppConfig

	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Warn("Cannot use env variables, use default")
		return &AppConfig{Profile: "dev"}
	}

	config.Profile = os.Getenv("PROFILE")

	return &config
}

func setupLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(r *runtime.Frame) (function string, file string) {
			filepath := strings.Split(r.File, "/")
			return "", fmt.Sprintf("%s:%v", filepath[len(filepath)-1], r.Line)
		},
	})

	if Config.Profile == "prod" {
		file, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logrus.SetOutput(os.Stdout)
			logrus.WithError(err).Warn("Cannot open log file. Logging to stdout.")
		} else {
			logrus.SetOutput(file)
		}
	} else {
		logrus.SetOutput(os.Stdout)
	}

	// adds information about location of log
	logrus.SetReportCaller(true)
}

func setupScraperModule() *scraper.Module {
	ceneoScraper := scraper.NewCeneoScraper()
	scraperModule, err := scraper.New(map[scraper.WebsiteType]scraper.WebsiteScraper{
		scraper.Ceneo: ceneoScraper,
	})
	if err != nil {
		logrus.WithError(err).Panic("Can't initialize search module.")
	}
	return scraperModule
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
	result, _ := scraperModule.Search(request)
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
	result, _ := scraperModule.CheckPrice(request)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func checkDatabaseHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: "ok"})
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
