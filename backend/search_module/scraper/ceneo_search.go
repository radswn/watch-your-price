package scraper

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func (cs *CeneoScraper) createSearchUrl(phrase string, page int) string {
	url := strings.Join([]string{cs.BaseUrl, phrase}, "")

	if page > 0 {
		url = strings.Join([]string{url, ";0020-30-0-0-", strconv.Itoa(page)}, "")
	}

	url = strings.Join([]string{url, ".htm?nocatnarrow=1"}, "")

	return url
}

func (cs *CeneoScraper) search(url string, phrase string, page int) (SearchResult, error) {
	result := SearchResult{
		Phrase:  phrase,
		Page:    page,
		Results: make(map[string]string),
	}

	c, err := cs.createSearchCollector(&result)
	if err != nil {
		logrus.WithError(err).Error("can't create collector")
		return SearchResult{}, err
	}

	q, err := cs.createQueue()
	if err != nil {
		logrus.WithError(err).Error("cannot create queue for ceneo")
		return SearchResult{}, err
	}

	err = q.AddURL(url)
	if err != nil {
		logrus.WithError(err).Error("error while adding url to search queue")
		return SearchResult{}, err
	}

	err = q.Run(c)
	if err != nil {
		logrus.WithError(err).Error("error while running collector")
		return SearchResult{}, err
	}

	c.Wait()

	return result, nil
}

func (cs *CeneoScraper) createSearchCollector(result *SearchResult) (*colly.Collector, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(cs.Domain),
	)

	err := cs.addLimitToCollector(c)
	if err != nil {
		logrus.WithError(err).Error("could not limit collector")
		return nil, err
	}

	cs.checkPageNumber(c, &result.NumOfPages)

	cs.handleItems(c, result.Results)

	return c, nil
}

func (cs *CeneoScraper) checkPageNumber(collector *colly.Collector, numOfPages *int) {
	collector.OnHTML("#page-counter", func(h *colly.HTMLElement) {
		number, err := strconv.Atoi(h.Attr("data-pagecount"))
		if err != nil {
			logrus.WithError(err).Warn("could not get number of pages")
			number = 0
		}
		*numOfPages = number
	})
}

type ceneoItemElement interface {
	getName() (string, error)
	getLink() (string, error)
	linkToAnotherShop() bool
}

func (cs *CeneoScraper) handleItems(collector *colly.Collector, results map[string]string) {
	collector.OnHTML("strong.cat-prod-row__name, div.grid-row", func(h *colly.HTMLElement) {

		var item ceneoItemElement
		if cs.isGridView(h) {
			item = ceneoGridItem{htmlElement: h}
		} else {
			item = ceneoListItem{htmlElement: h}
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

func (cs *CeneoScraper) isGridView(h *colly.HTMLElement) bool {
	return strings.EqualFold(h.Name, "div")
}

type ceneoGridItem struct {
	htmlElement *colly.HTMLElement
}

func (cgi ceneoGridItem) getName() (string, error) {
	name := cgi.getLinkTag().SiblingsFiltered("div.grid-item__caption").Find("strong").First().Text()
	name = strings.TrimSpace(name)
	if strings.EqualFold("", name) {
		return "", errors.New("name attribute does not exist")
	}
	return name, nil
}

func (cgi ceneoGridItem) getLink() (string, error) {
	relativeLink, exists := cgi.getLinkTag().Attr("href")
	if !exists {
		return "", errors.New("href attribute does not exist")
	}
	link := cgi.htmlElement.Request.AbsoluteURL(relativeLink)
	return link, nil
}

func (cgi ceneoGridItem) getLinkTag() *goquery.Selection {
	return cgi.htmlElement.DOM.Find("a").First()
}

func (cgi ceneoGridItem) linkToAnotherShop() bool {
	return cgi.getLinkTag().HasClass("go-to-shop")
}

type ceneoListItem struct {
	htmlElement *colly.HTMLElement
}

func (cli ceneoListItem) getName() (string, error) {
	name := strings.TrimSpace(cli.getLinkTag().Text())
	if strings.EqualFold("", name) {
		return "", errors.New("name attribute does not exist")
	}
	return name, nil
}

func (cli ceneoListItem) getLink() (string, error) {
	relativeLink, exists := cli.getLinkTag().Attr("href")
	if !exists {
		return "", errors.New("href attribute not exists")
	}
	link := cli.htmlElement.Request.AbsoluteURL(relativeLink)
	return link, nil
}

func (cli ceneoListItem) getLinkTag() *goquery.Selection {
	return cli.htmlElement.DOM.Find("a").First()
}

func (cli ceneoListItem) linkToAnotherShop() bool {
	return cli.getLinkTag().HasClass("go-to-shop")
}
