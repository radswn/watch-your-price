package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"search_module/search"
	"search_module/search/website_type"
	"search_module/search/websites"
	"strings"
)

var searchModule *search.SearchModule

func init() {
	setupLogrus()
	searchModule = setupSearchModule()
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
