package ceneo

import (
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
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

	q, _ := queue.New(
		4,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	c.OnHTML("a.js_pagination-top-next", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		q.AddURL(h.Request.AbsoluteURL(link))
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

	q.AddURL(url)
	q.Run(c)

	c.Wait()

	return results
}

//TODO
// func checkPrice(url string) string {

// 	c.OnHTML("h1.js_product-h1-link", func(h *colly.HTMLElement) {
// 		h.DOM.ParentsUntil("~").Find("meta").Each(func(_ int, s *goquery.Selection) {
// 			property, _ := s.Attr("property")
// 			if strings.EqualFold(property, "product:price:amount") {
// 				price, _ := s.Attr("content")
// 				fmt.Println(strings.TrimSpace(h.Text) + "|" + price)
// 			}
// 		})
// 	})
// }
