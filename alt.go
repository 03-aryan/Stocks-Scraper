package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StockData struct {
	Symbol string `json:"01. symbol"`
	Price  string `json:"05. price"`
}

func main() {
	apiKey := "YOUR_ALPHA_VANTAGE_API_KEY"
	ticker := []string{"IBM", "AAPL", "GOOGL"}

	for _, t := range ticker {
		url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", t, apiKey)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		stock := result["Global Quote"].(map[string]interface{})
		fmt.Printf("Symbol: %s, Price: %s\n", stock["01. symbol"], stock["05. price"])
	}
}
