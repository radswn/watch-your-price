package websites

import (
	"backend/search_module"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type ceneoSearch struct {
}

const ceneoUrl = "https://www.ceneo.pl/;szukaj-"

func (cs *ceneoSearch) GetResults(phrase string, page int) (search_module.SearchResult, error) {

	url := createSearchUrl(phrase, page)

	c := colly.NewCollector(
		colly.AllowedDomains("www.ceneo.pl"),
	)

	q, _ := queue.New(
		4,
		&queue.InMemoryQueueStorage{MaxSize: 100},
	)

	err := addLimitToCollector(c)
	if err != nil {
		logrus.WithError(err).Error("could not limit collector")
		return search_module.SearchResult{}, err
	}

	var maxPages int
	checkPageNumber(c, &maxPages)

	results := make(map[string]string)
	handleItemsOnGridView(c, results)
	handleItemsOnListView(c, results)

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

	return search_module.SearchResult{
		Phrase:     phrase,
		Page:       page,
		NumOfPages: maxPages,
		Results:    results,
	}, nil
}

func createSearchUrl(phrase string, page int) string {
	url := strings.Join([]string{ceneoUrl, phrase}, "")
	if page > 0 {
		url = url + ";0020-30-0-0-" + strconv.Itoa(page) + ".htm"
		url = strings.Join([]string{url, ";0020-30-0-0-", strconv.Itoa(page), ".htm"}, "")
	}
	return url
}

func addLimitToCollector(collector *colly.Collector) error {
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "www.ceneo.pl/*",
		Delay:       3 * time.Second,
		RandomDelay: 1 * time.Second,
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

func handleItemsOnGridView(collector *colly.Collector, results map[string]string) {
	collector.OnHTML("div.grid-row", func(h *colly.HTMLElement) {
		linkTag := h.DOM.Find("a").First()
		if linkTag.HasClass("go-to-shop") {
			return
		}
		relativeLink, _ := linkTag.Attr("href")
		link := h.Request.AbsoluteURL(relativeLink)
		name := linkTag.SiblingsFiltered("div.grid-item__caption").Find("Strong").First().Text()
		results[strings.TrimSpace(name)] = link
	})
}

func handleItemsOnListView(collector *colly.Collector, results map[string]string) {
	collector.OnHTML("strong.cat-prod-row__name", func(h *colly.HTMLElement) {
		linkTag := h.DOM.Find("a").First()
		if linkTag.HasClass("go-to-shop") {
			return
		}
		relativeLink, _ := linkTag.Attr("href")
		link := h.Request.AbsoluteURL(relativeLink)
		results[strings.TrimSpace(linkTag.Text())] = link
	})
}
