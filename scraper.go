package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

var Items []item


func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("h2.product-title"),
			Price: h.ChildText("div.sale-price"),
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

	c.Visit("https://j2store.net/demo/index.php/shop")
	//fmt.Println(Items)

	//Creating a csvFile
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()
     // Writing to a csvFile
	writer := csv.NewWriter(file)
	for _, element :=  range Items {
		data := []string{element.Name, element.Price, element.ImageUrl}
		err := writer.Write(data)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	writer.Flush()
}

