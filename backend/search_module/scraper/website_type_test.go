package scraper_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"search_module/scraper"
	"testing"
)

func TestUnmarshallingWebsiteTypeWithCorrectValueShouldReturnEnumWithAppropriateType(t *testing.T) {
	jsonInput := []byte(`"ceneo"`)
	var wt scraper.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.Nil(t, err)
	assert.Equal(t, wt, scraper.Ceneo)
}

func TestUnmarshallingWebsiteTypeWithEmptyValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`""`)
	var wt scraper.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}

func TestUnmarshallingWebsiteTypeWithNonExistingValueShouldReturnError(t *testing.T) {
	jsonInput := []byte(`"not_exist"`)
	var wt scraper.WebsiteType

	err := json.Unmarshal(jsonInput, &wt)

	assert.NotNil(t, err)
}
