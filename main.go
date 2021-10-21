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
		symbol := e.ChildText(".d-lg-none")
		market_cap := e.ChildText(".td-market_cap")
		price := e.ChildText(".td-price")
		volume_24h := e.ChildText(".td-liquidity_score")
		percent_change_1h := e.ChildText(".td-change1h")
		percent_change_24h := e.ChildText(".td-change24h")
		percent_change_7d := e.ChildText(".td-change7d")

		coin := Coin{
			Symbol:           symbol,
			MarketCap:        market_cap,
			Price:            price,
			Volume24H:        volume_24h,
			PercentChange1H:  percent_change_1h,
			PercentChange24H: percent_change_24h,
			PercentChange7D:  percent_change_7d,
		}
	})

	c.Visit("https://www.coingecko.com/en")

}
