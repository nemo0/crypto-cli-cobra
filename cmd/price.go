/*
Copyright © 2022 Subha Chanda <subhachanda88@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [token]",
	Short: "Get price information of a token",
	Long: `Get price information of a token. For example: crypto-cli price bitcoin will return the price of bitcoin. 
	To get the markets for a token, use the --markets flag. For example: crypto-cli price bitcoin --markets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currency := strings.ToLower(args[0])
		showMarketData, _ := cmd.Flags().GetBool("markets")

		if showMarketData {
			getMarkets(currency)
		} else {
			getCurrencyPrice(currency)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


type Response struct {
	Data CryptoData `json:"data"`
}

type CryptoData struct {
	Id string `json:"id"`
	Rank string `json:"rank"`
	Symbol string `json:"symbol"`
	Name string `json:"name"`
	Supply string `json:"supply"`
	MaxSupply string `json:"maxSupply"`
	MarketCapUsd string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr string `json:"vwap24Hr"`
}

func getCurrencyPrice(currency string) {
	fmt.Println("Getting price for ", currency)
	coincapApiUrl := "https://api.coincap.io/v2/assets/" + currency

	client := http.Client{}

	req, err := http.NewRequest("GET", coincapApiUrl, nil)

	if err != nil {
		log.Fatal(err, "Error creating request")
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err, "Error sending request")
	}

	defer res.Body.Close()

  	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err, "Error reading response body")
	}

	var data Response
	json.Unmarshal(respBody, &data)



	if data.Data.Id == "" {
		fmt.Println("Currency not found")
	} else {
		fmt.Printf("Currency: %s\n Symbol: %s\n Price: $%s\n Market Cap: $%s\n Volume: %s\n Change: %s%%\n Rank: %s\n Supply: %s\n Max Supply: %s\n Vwap: %s\n", data.Data.Name, data.Data.Symbol, data.Data.PriceUsd, data.Data.MarketCapUsd, data.Data.VolumeUsd24Hr, data.Data.ChangePercent24Hr, data.Data.Rank, data.Data.Supply, data.Data.MaxSupply, data.Data.Vwap24Hr)
	}
}


type MarketResponse struct {
	Data []MarketData `json:"data"`
}

type MarketData struct {
	ExchangeId string `json:"exchangeId"`
	BaseId string `json:"baseId"`
	QuoteId string `json:"quoteId"`
	QuoteSymbol string `json:"quoteSymbol"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd string `json:"priceUsd"`
	VolumePercent string `json:"volumePercent"`
}

func getMarkets(currency string) {
	fmt.Println("Getting markets for ", currency)
	coincapApiUrl := "https://api.coincap.io/v2/assets/" + currency + "/markets?limit=20"

	client := http.Client{}

	req, err := http.NewRequest("GET", coincapApiUrl, nil)

	if err != nil {
		log.Fatal(err, "Error creating request")
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err, "Error sending request")
	}

	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err, "Error reading response body")
	}

	var data MarketResponse
	json.Unmarshal(respBody, &data)


	for _, market := range data.Data {
		fmt.Printf("\n\nExchange: %s\nBase: %s\nQuote: %s\nPrice: $%s\nVolume: %s\nVolume Percent: %s\n", market.ExchangeId, market.BaseId, market.QuoteSymbol, market.PriceUsd, market.VolumeUsd24Hr, market.VolumePercent)
	}
}