package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func btcToUah() decimal.Decimal {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?start=1&limit=5000&convert=USD", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	q.Add("convert", "UAH")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "9df5fa6e-da70-4df1-baab-ed5cd829b5b2")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	var btc BTCdata
	err = json.Unmarshal(respBody, &btc)
	if err != nil {
		fmt.Println("error:", err)
	}

	return btc.Data[0].Quote.Uah.Price
}

type BTCdata struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
		TotalCount   int         `json:"total_count"`
	} `json:"status"`
	Data []struct {
		ID                            int         `json:"id"`
		Name                          string      `json:"name"`
		Symbol                        string      `json:"symbol"`
		Slug                          string      `json:"slug"`
		NumMarketPairs                int         `json:"num_market_pairs"`
		DateAdded                     time.Time   `json:"date_added"`
		Tags                          []string    `json:"tags"`
		MaxSupply                     int         `json:"max_supply"`
		CirculatingSupply             int         `json:"circulating_supply"`
		TotalSupply                   int         `json:"total_supply"`
		Platform                      interface{} `json:"platform"`
		CmcRank                       int         `json:"cmc_rank"`
		SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
		SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
		TvlRatio                      interface{} `json:"tvl_ratio"`
		LastUpdated                   time.Time   `json:"last_updated"`
		Quote                         struct {
			Uah struct {
				Price                 decimal.Decimal `json:"price"`
				Volume24H             float64         `json:"volume_24h"`
				VolumeChange24H       float64         `json:"volume_change_24h"`
				PercentChange1H       float64         `json:"percent_change_1h"`
				PercentChange24H      float64         `json:"percent_change_24h"`
				PercentChange7D       float64         `json:"percent_change_7d"`
				PercentChange30D      float64         `json:"percent_change_30d"`
				PercentChange60D      float64         `json:"percent_change_60d"`
				PercentChange90D      float64         `json:"percent_change_90d"`
				MarketCap             float64         `json:"market_cap"`
				MarketCapDominance    float64         `json:"market_cap_dominance"`
				FullyDilutedMarketCap float64         `json:"fully_diluted_market_cap"`
				Tvl                   interface{}     `json:"tvl"`
				LastUpdated           time.Time       `json:"last_updated"`
			} `json:"UAH"`
		} `json:"quote"`
	} `json:"data"`
}
