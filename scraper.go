package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var Items []item

type item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name:     h.ChildText("h2.product-title"),
			Price:    h.ChildText("div.sale-price"),
			ImageUrl: h.ChildAttr("img", "src"),
		}
		Items = append(Items, item)
	})


	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})


	c.Visit("http://j2store.net/demo/index.php/shop")


	content, err := json.Marshal(Items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("json-Content", content, 0644)

}

