package ceneo

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func searchForItem(name string) map[string]string {
	results := make(map[string]string)
	url := "https://www.ceneo.pl/;szukaj-" + name
	c := colly.NewCollector(
		colly.AllowedDomains("www.ceneo.pl"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "www.ceneo.pl/*",
		Delay:       10 * time.Second,
		RandomDelay: 10 * time.Second,
	})

	c.OnHTML("a.js_pagination-top-next", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
	})

	c.OnHTML("a.grid-item__thumb[href]", func(h *colly.HTMLElement) {
		if h.DOM.HasClass("go-to-shop") {
			return
		}
		link := h.Request.AbsoluteURL(h.Attr("href"))
		name = h.DOM.SiblingsFiltered("div.grid-item__caption").Find("Strong").First().Text()
		results[name] = link
	})

	c.OnHTML("a.go-to-product[href]", func(h *colly.HTMLElement) {
		if h.DOM.HasClass("go-to-shop") {
			return
		}
		link := h.Request.AbsoluteURL(h.Attr("href"))
		results[h.Text] = link
	})

	c.Visit(url)
	return results
}
