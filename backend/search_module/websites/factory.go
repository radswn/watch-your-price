package websites

import (
	"backend/search_module"
	"backend/search_module/website_type"
	"github.com/sirupsen/logrus"
)

// New returns new instance of SearchModule with provided websites
func New(websiteType website_type.WebsiteType) (search_module.WebsiteSearch, error) {
	var module search_module.WebsiteSearch
	switch websiteType {
	case website_type.Ceneo:
		queue, err := createQueue()
		if err != nil {
			logrus.WithError(err).Error("cannot create queue for ceneo")
			return nil, err
		}
		module = &ceneoSearch{queue: queue}
	}
	return module, nil
}
