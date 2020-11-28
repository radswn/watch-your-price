package search_module_test

import (
	"backend/search_module"
	"backend/search_module/website_type"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestUnmarshallingWebsiteTypeWithCorrectValueShouldReturnEnumWithAppropriateType(t *testing.T) {
	jsonInput := []byte(`"ceneo"`)
	var wt website_type.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.Nil(t, err)
	assert.Equal(t, wt, website_type.Ceneo)
}

func TestUnmarshallingWebsiteTypeWithEmptyValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`""`)
	var wt website_type.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}

func TestUnmarshallingWebsiteTypeWithNonExistingValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`"not_exist"`)
	var wt website_type.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}
