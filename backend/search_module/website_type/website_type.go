package website_type

import (
	"encoding/json"
	"errors"
)

type WebsiteType string

const (
	Ceneo WebsiteType = "ceneo"
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
	}
	return errors.New("Invalid website type")
}
