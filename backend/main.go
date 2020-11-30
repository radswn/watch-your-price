package main

import (
	"backend/search_module"
	"backend/search_module/website_type"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	file, err := os.OpenFile("backend.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.SetOutput(os.Stdout)
		logrus.WithError(err).Warn("Cannot open log file. Logging to stdout.")
	} else {
		logrus.SetOutput(file)
	}

	// adds information about location of log
	logrus.SetReportCaller(true)
}

var db = make(map[string]string)
var searchModule = setupSearchModule()

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
	r.POST("/search", func(c *gin.Context) {
		var request search_module.SearchRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logrus.WithError(err).Info("Could not parse search request.")
			return
		}
		result, err := searchModule.Search(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			logrus.WithError(err).Error("Could not get search results.")
			return
		}
		c.JSON(http.StatusOK, result)
	})
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

// use this implementation of website search before I port ceneo module - Dawid
type testWebsiteSearch struct {
}

func (tws testWebsiteSearch) GetResults(phrase string, page int) (search_module.SearchResult, error) {
	sr := search_module.SearchResult{
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

func setupSearchModule() *search_module.SearchModule {
	searchModule, err := search_module.New(map[website_type.WebsiteType]search_module.WebsiteSearch{
		website_type.Ceneo: testWebsiteSearch{},
	})
	if err != nil {
		logrus.WithError(err).Panic("Can't initialize search module.")
	}
	return searchModule
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
