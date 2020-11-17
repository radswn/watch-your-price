package ceneo

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	url := "https://www.ceneo.pl/SiteMap.aspx"
	// fi, _ := os.Create("file.txt")

	c := colly.NewCollector(
		colly.AllowedDomains("www.ceneo.pl"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "www.ceneo.pl/*",
		Delay:       10 * time.Second,
		RandomDelay: 10 * time.Second,
	})

	q, _ := queue.New(
		4,
		&queue.InMemoryQueueStorage{MaxSize: 500000},
	)

	c.OnHTML("a.js_pagination-top-next", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		q.AddURL(link)
	})

	c.OnHTML("a.grid-item__thumb[href]", func(h *colly.HTMLElement) {
		if h.DOM.HasClass("go-to-shop") {
			return
		}
		link := h.Request.AbsoluteURL(h.Attr("href"))
		q.AddURL(link)
	})

	c.OnHTML("a.go-to-product[href]", func(h *colly.HTMLElement) {
		if h.DOM.HasClass("go-to-shop") {
			return
		}
		link := h.Request.AbsoluteURL(h.Attr("href"))
		q.AddURL(link)
	})

	c.OnHTML("dl.category-links", func(h *colly.HTMLElement) {

		// name, _ := h.DOM.ChildrenFiltered("dt").ChildrenFiltered("a").Attr("href")

		// if strings.EqualFold(name, "/Bizuteria_i_zegarki") {

		h.DOM.Find("li").Not("li.has-children").Find("a").Each(func(_ int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			q.AddURL(h.Request.AbsoluteURL(href))
		})
		// }

	})

	c.OnHTML("h1.js_product-h1-link", func(h *colly.HTMLElement) {
		h.DOM.ParentsUntil("~").Find("meta").Each(func(_ int, s *goquery.Selection) {
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "product:price:amount") {
				price, _ := s.Attr("content")
				// fmt.Fprintln(fi, strings.TrimSpace(h.Text)+"|"+price)
				fmt.Println(strings.TrimSpace(h.Text) + "|" + price)
			}
		})
	})

	q.AddURL(url)
	q.Run(c)
}
