package check

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type ceneoCheck struct {
	domain      string
	domainGlob  string
	delay       time.Duration
	randomDelay time.Duration
}

func newCeneoCheck() *ceneoCheck {
	return &ceneoCheck{
		domain:      "www.ceneo.pl",
		domainGlob:  "www.ceneo.pl/*",
		delay:       3 * time.Second,
		randomDelay: 1 * time.Second,
	}
}

func (cc *ceneoCheck) GetResults(url string) (CheckResult, error) {

	result, err := cc.check(url)
	if err != nil {
		logrus.WithError(err).Error("can't process check request")
		return CheckResult{}, err
	}

	return result, nil
}

func (cc *ceneoCheck) check(url string) (CheckResult, error) {
	var price string

	c := colly.NewCollector(
		colly.AllowedDomains(cc.domain),
	)

	err := cc.addLimitToCollector(c)
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

func (cc *ceneoCheck) addLimitToCollector(collector *colly.Collector) error {
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  cc.domainGlob,
		Delay:       cc.delay,
		RandomDelay: cc.randomDelay,
	})
	return err
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
