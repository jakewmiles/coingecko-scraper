package main

import (
	"github.com/gocolly/colly/v2"
)

type Coin struct {
	Symbol           string
	MarketCap        string
	Price            string
	Volume24H        string
	PercentChange1H  string
	PercentChange24H string
	PercentChange7D  string
}

func main() {

	c := colly.NewCollector()

	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {

	})

	c.Visit("https://www.coingecko.com/en")

}
