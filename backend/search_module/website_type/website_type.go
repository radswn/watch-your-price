package website_type

import (
	"encoding/json"
	"errors"
)

type WebsiteType string

const (
	Ceneo WebsiteType = "ceneo"
	Empty WebsiteType = ""
)

func (wt *WebsiteType) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type WT WebsiteType
	var r *WT = (*WT)(wt)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *wt {
	case Ceneo:
		return nil
	case Empty:
		return errors.New("Empty value")
	}
	return errors.New("Invalid leave type")
}
