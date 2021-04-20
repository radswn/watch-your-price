package scraper

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/sirupsen/logrus"
	"time"
)

type CeneoScraper struct {
	domain       string
	domainGlob   string
	delay        time.Duration
	randomDelay  time.Duration
	queueStorage int
	queueThreads int
	baseUrl      string
}

func NewCeneoScraper() *CeneoScraper {
	return &CeneoScraper{
		queueStorage: 100,
		queueThreads: 4,
		domain:       "www.ceneo.pl",
		domainGlob:   "www.ceneo.pl/*",
		baseUrl:      "https://www.ceneo.pl/;szukaj-",
		delay:        3 * time.Second,
		randomDelay:  1 * time.Second,
	}
}

func (cs *CeneoScraper) Search(phrase string, page int) (SearchResult, error) {

	url := cs.createSearchUrl(phrase, page)

	result, err := cs.search(url, phrase, page)
	if err != nil {
		logrus.WithError(err).Error("can't process search request")
		return SearchResult{}, err
	}

	return result, nil
}

func (cs *CeneoScraper) CheckPrice(url string) (CheckResult, error) {

	result, err := cs.check(url)
	if err != nil {
		logrus.WithError(err).Error("can't process check request")
		return CheckResult{}, err
	}

	return result, nil
}

func (cs *CeneoScraper) addLimitToCollector(collector *colly.Collector) error {
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  cs.domainGlob,
		Delay:       cs.delay,
		RandomDelay: cs.randomDelay,
	})
	return err
}

func (cs *CeneoScraper) createQueue() (*queue.Queue, error) {
	q, err := queue.New(
		cs.queueThreads,
		&queue.InMemoryQueueStorage{MaxSize: cs.queueStorage},
	)

	if err != nil {
		logrus.WithError(err).Error("can not create ceneo queue")
		return nil, err
	}

	return q, nil
}
