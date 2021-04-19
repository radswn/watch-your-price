package check_test

import (
	"search_module/scraper"
	"search_module/scraper/check"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

type testWebsiteCheck struct {
}

func (twc testWebsiteCheck) GetResults(url string) (check.Result, error) {
	result := check.Result{
		Price: "10",
	}
	return result, nil
}

func TestCheckShouldReturnResultsFromWebsiteCheckImplementation(t *testing.T) {
	websiteCheckMap := make(map[scraper.WebsiteType]check.WebsiteCheck)
	websiteCheckMap[scraper.Ceneo] = testWebsiteCheck{}
	module, err := check.NewCheck(websiteCheckMap)
	assert.Nil(t, err)

	requestData := check.Request{
		Url:     "example.com/3",
		Website: "ceneo",
	}

	expectedPrice := "10"

	result, err := module.Check(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPrice, result.Price)
}
