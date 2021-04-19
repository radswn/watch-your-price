package search_test

import (
	"encoding/json"
	"search_module/search"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

type testWebsiteSearch struct {
}

func (tws testWebsiteSearch) GetResults(phrase string, page int) (search.Result, error) {
	sr := search.Result{
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

func TestUnmarshallingWebsiteTypeWithCorrectValueShouldReturnEnumWithAppropriateType(t *testing.T) {
	jsonInput := []byte(`"ceneo"`)
	var wt search.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.Nil(t, err)
	assert.Equal(t, wt, search.Ceneo)
}

func TestUnmarshallingWebsiteTypeWithEmptyValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`""`)
	var wt search.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}

func TestUnmarshallingWebsiteTypeWithNonExistingValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`"not_exist"`)
	var wt search.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}

func TestSearchShouldReturnResultsFromWebsiteSearchImplementation(t *testing.T) {
	websiteSearchMap := make(map[search.WebsiteType]search.WebsiteSearch)
	websiteSearchMap[search.Ceneo] = testWebsiteSearch{}
	module, err := search.New(websiteSearchMap)
	assert.Nil(t, err)

	requestData := search.Request{
		Phrase:  "test",
		Page:    3,
		Website: "ceneo",
	}

	expectedPhrase := "test"
	expectedPage := 3
	expectedResults := map[string]string{
		"result1": "example.com/1",
		"result2": "example.com/2",
		"result3": "example.com/3",
		"result4": "example.com/4",
	}

	result, err := module.Search(requestData)

	assert.Nil(t, err)
	assert.Equal(t, expectedPhrase, result.Phrase)
	assert.Equal(t, expectedPage, result.Page)
	assert.Equal(t, expectedResults, result.Results)
}
