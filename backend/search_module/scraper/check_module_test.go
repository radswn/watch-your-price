package scrapper_test

import (
	"search_module/scrapper"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

type testWebsiteCheck struct {
}

func (twc testWebsiteCheck) GetResults(url string) (scrapper.Result, error) {
	result := scrapper.CheckResult{
		Price: 10,
	}
	return result, nil
}

func TestCheckShouldReturnResultsFromWebsiteCheckImplementation(t *testing.T) {
	websiteCheckMap := make(map[scrapper.WebsiteType]scrapper.WebsiteCheck)
	websiteCheckMap[scrapper.Ceneo] = testWebsiteCheck{}
	module, err := scrapper.New(websiteCheckMap)
	assert.Nil(t, err)

	requestData := scrapper.CheckRequest{
		Url:     "example.com/3",
		Website: "ceneo",
	}

	expectedPrice := 10

	result, err := module.Check(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPrice, result.Price)
}
