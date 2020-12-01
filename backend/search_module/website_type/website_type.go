package website_type

import (
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
)

// WebsiteType represent enum with values of implemented search engines
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
		logrus.WithError(err).Info("Wrong input type for parameter website.")
		return errors.New("Wrong value type for parameter website")
	}

	if isValueInWebsiteTypeEnum(wt) {
		return nil
	}
	err = errors.New("Invalid website type")
	logrus.WithError(err).Info()
	return err
}

func isValueInWebsiteTypeEnum(value *WebsiteType) bool {
	switch *value {
	case Ceneo:
		return true
	}
	return false
}
