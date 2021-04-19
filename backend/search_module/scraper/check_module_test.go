package scraper_test

import (
	"search_module/scraper"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

type testWebsiteCheck struct {
}

func (twc testWebsiteCheck) GetResults(url string) (scraper.CheckResult, error) {
	result := scraper.CheckResult{
		Price: "10",
	}
	return result, nil
}

func TestCheckShouldReturnResultsFromWebsiteCheckImplementation(t *testing.T) {
	websiteCheckMap := make(map[scraper.WebsiteType]scraper.WebsiteCheck)
	websiteCheckMap[scraper.Ceneo] = testWebsiteCheck{}
	module, err := scraper.NewCheck(websiteCheckMap)
	assert.Nil(t, err)

	requestData := scraper.CheckRequest{
		Url:     "example.com/3",
		Website: "ceneo",
	}

	expectedPrice := "10"

	result, err := module.Check(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPrice, result.Price)
}
