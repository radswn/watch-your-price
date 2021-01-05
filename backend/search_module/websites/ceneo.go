package websites

import (
	"backend/search_module"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"strings"
	"time"
)

type ceneoSearch struct {
}

// New returns new instance of SearchModule with provided websites
func New() (*ceneoSearch, error) {
	search := &ceneoSearch{}
	return search, nil
}

func (cs *ceneoSearch) GetResults(phrase string, page int) (search_module.SearchResult, error) {

	searchResult := search_module.SearchResult{Phrase: phrase, Page: page}
	searchResult.NumOfPages = page + 1 //TODO implement num of pages checking
	results := make(map[string]string)

	url := "https://www.ceneo.pl/;szukaj-" + phrase

	c := colly.NewCollector(
		colly.AllowedDomains("www.ceneo.pl"),
	)

	q, _ := queue.New(
		4,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	addLimitToCollector(c)

	handleItemsOnGridView(c, results)

	handleItemsOnListView(c, results)

	q.AddURL(url)
	q.Run(c)

	c.Wait()
	searchResult.Results = results
	return searchResult, nil
}

func addLimitToCollector(collector *colly.Collector) {
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "www.ceneo.pl/*",
		Delay:       3 * time.Second,
		RandomDelay: 1 * time.Second,
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
