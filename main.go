package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	ticker := []string{
		"MSFT", "IBM", "GE", "UNP", "COST", "MCD", "V", "WMT", "DIS", "MMM",
		"INTC", "AXP", "AAPL", "BA", "CSCO", "GS", "JPM", "CRM", "VZ",
	}

	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.company = e.ChildText("h1")
		stock.price = e.ChildText(".D\\(ib\\) .Fz\\(36px\\)")
		stock.change = e.ChildText(".Trsdu\\(0\\.3s\\)")

		if stock.company != "" {
			stocks = append(stocks, stock)
		}
	})

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t)
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}

	c.Wait()

	if len(stocks) == 0 {
		log.Println("No stocks data collected.")
		return
	}

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatal("Could not create csv file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"company", "price", "change"}
	writer.Write(headers)

	for _, stock := range stocks {
		record := []string{stock.company, stock.price, stock.change}
		writer.Write(record)
	}

	fmt.Println("Data written to stocks.csv successfully!")
}
