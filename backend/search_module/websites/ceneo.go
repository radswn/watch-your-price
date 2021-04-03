package websites

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/sirupsen/logrus"
	"search_module"
	"strconv"
	"strings"
	"time"
)

type ceneoSearch struct {
	queueStorage int
	queueThreads int
	domain       string
	domainGlob   string
	baseUrl      string
	delay        time.Duration
	randomDelay  time.Duration
}

func newCeneoSearch() *ceneoSearch {
	return &ceneoSearch{
		queueStorage: 100,
		queueThreads: 4,
		domain:       "www.ceneo.pl",
		domainGlob:   "www.ceneo.pl/*",
		baseUrl:      "https://www.ceneo.pl/;szukaj-",
		delay:        3 * time.Second,
		randomDelay:  1 * time.Second,
	}
}

func (cs *ceneoSearch) GetResults(phrase string, page int) (search_module.SearchResult, error) {

	url := cs.createSearchUrl(phrase, page)

	result, err := cs.search(url, phrase, page)
	if err != nil {
		logrus.WithError(err).Error("can't process search request")
		return search_module.SearchResult{}, err
	}

	return result, nil
}

func (cs *ceneoSearch) createSearchUrl(phrase string, page int) string {
	url := strings.Join([]string{cs.baseUrl, phrase}, "")

	if page > 0 {
		url = strings.Join([]string{url, ";0020-30-0-0-", strconv.Itoa(page)}, "")
	}

	url = strings.Join([]string{url, ".htm?nocatnarrow=1"}, "")

	return url
}

func (cs *ceneoSearch) search(url string, phrase string, page int) (search_module.SearchResult, error) {
	result := search_module.SearchResult{
		Phrase:  phrase,
		Page:    page,
		Results: make(map[string]string),
	}

	c, err := cs.createCollector(&result)
	if err != nil {
		logrus.WithError(err).Error("can't create collector")
		return search_module.SearchResult{}, err
	}

	q, err := cs.createQueue()
	if err != nil {
		logrus.WithError(err).Error("cannot create queue for ceneo")
		return search_module.SearchResult{}, err
	}

	err = q.AddURL(url)
	if err != nil {
		logrus.WithError(err).Error("error while adding url to search queue")
		return search_module.SearchResult{}, err
	}

	err = q.Run(c)
	if err != nil {
		logrus.WithError(err).Error("error while running collector")
		return search_module.SearchResult{}, err
	}

	c.Wait()

	return result, nil
}

func (cs *ceneoSearch) createCollector(result *search_module.SearchResult) (*colly.Collector, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(cs.domain),
	)

	err := cs.addLimitToCollector(c)
	if err != nil {
		logrus.WithError(err).Error("could not limit collector")
		return nil, err
	}

	checkPageNumber(c, &result.NumOfPages)

	handleItems(c, result.Results)

	return c, nil
}

func (cs *ceneoSearch) createQueue() (*queue.Queue, error) {
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

func (cs *ceneoSearch) addLimitToCollector(collector *colly.Collector) error {
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  cs.domainGlob,
		Delay:       cs.delay,
		RandomDelay: cs.randomDelay,
	})
	return err
}

func checkPageNumber(collector *colly.Collector, numOfPages *int) {
	collector.OnHTML("#page-counter", func(h *colly.HTMLElement) {
		number, err := strconv.Atoi(h.Attr("data-pagecount"))
		if err != nil {
			logrus.WithError(err).Warn("could not get number of pages")
			number = 0
		}
		*numOfPages = number
	})
}

type itemElement interface {
	getName() (string, error)
	getLink() (string, error)
	linkToAnotherShop() bool
}

func handleItems(collector *colly.Collector, results map[string]string) {
	collector.OnHTML("strong.cat-prod-row__name, div.grid-row", func(h *colly.HTMLElement) {

		var item itemElement
		if isGridView(h) {
			item = gridItem{htmlElement: h}
		} else {
			item = listItem{htmlElement: h}
		}

		if item.linkToAnotherShop() {
			return
		}
		name, err := item.getName()
		if err != nil {
			logrus.WithError(err).Warn("could not find name")
			return
		}
		link, err := item.getLink()
		if err != nil {
			logrus.WithError(err).Warn("could not find link for " + name)
			return
		}
		results[name] = link
	})
}

func isGridView(h *colly.HTMLElement) bool {
	return strings.EqualFold(h.Name, "div")
}

type gridItem struct {
	htmlElement *colly.HTMLElement
}

func (gi gridItem) getName() (string, error) {
	name := gi.getLinkTag().SiblingsFiltered("div.grid-item__caption").Find("strong").First().Text()
	name = strings.TrimSpace(name)
	if strings.EqualFold("", name) {
		return "", errors.New("name attribute does not exist")
	}
	return name, nil
}

func (gi gridItem) getLink() (string, error) {
	relativeLink, exists := gi.getLinkTag().Attr("href")
	if !exists {
		return "", errors.New("href attribute does not exist")
	}
	link := gi.htmlElement.Request.AbsoluteURL(relativeLink)
	return link, nil
}

func (gi gridItem) getLinkTag() *goquery.Selection {
	return gi.htmlElement.DOM.Find("a").First()
}

func (gi gridItem) linkToAnotherShop() bool {
	return gi.getLinkTag().HasClass("go-to-shop")
}

type listItem struct {
	htmlElement *colly.HTMLElement
}

func (li listItem) getName() (string, error) {
	name := strings.TrimSpace(li.getLinkTag().Text())
	if strings.EqualFold("", name) {
		return "", errors.New("name attribute does not exist")
	}
	return name, nil
}

func (li listItem) getLink() (string, error) {
	relativeLink, exists := li.getLinkTag().Attr("href")
	if !exists {
		return "", errors.New("href attribute not exists")
	}
	link := li.htmlElement.Request.AbsoluteURL(relativeLink)
	return link, nil
}

func (li listItem) getLinkTag() *goquery.Selection {
	return li.htmlElement.DOM.Find("a").First()
}

func (li listItem) linkToAnotherShop() bool {
	return li.getLinkTag().HasClass("go-to-shop")
}
