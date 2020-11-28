package search_module_test

import (
	"backend/search_module"
	"backend/search_module/website_type"
	"encoding/json"
	"testing"
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

func TestUnmarshallingJsonRequestShouldReturnObjectWithCorrectFields(t *testing.T) {
	jsonInput := []byte(`{"phrase": "test", "page": 3, "Website": "ceneo"}`)
	var request search_module.SearchRequest
	err := json.Unmarshal(jsonInput, &request)
	if err != nil {
		t.Fatal(err)
	}
	if request.Phrase != "test" {
		t.Fatal("Incorrect phrase in request object")
	}

	if request.Page != 3 {
		t.Fatal("Incorrect page number in request object")
	}

	if request.Website != website_type.Ceneo {
		t.Fatal("Incorrect website type in request object")
	}
}
