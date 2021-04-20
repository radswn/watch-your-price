package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func (cs *CeneoScraper) check(url string) (CheckResult, error) {
	var price string

	c := colly.NewCollector(
		colly.AllowedDomains(cs.domain),
	)

	err := cs.addLimitToCollector(c)
	if err != nil {
		logrus.WithError(err).Error("error while limiting collector")
		return CheckResult{}, err
	}

	findPriceTagOnPage(c, &price)

	err = c.Visit(url)
	if err != nil {
		logrus.WithError(err).Error("error while running collector")
		return CheckResult{}, err
	}
	return CheckResult{Price: price}, nil
}

func findPriceTagOnPage(collector *colly.Collector, price *string) {

	collector.OnHTML("html", func(h *colly.HTMLElement) {
		h.DOM.Find("meta").Each(func(_ int, s *goquery.Selection) {
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "product:price:amount") {
				*price, _ = s.Attr("content")
			}
		})
	})
}
