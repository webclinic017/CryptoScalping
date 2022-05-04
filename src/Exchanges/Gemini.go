package Exchanges

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func GetGeminiOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	/*
		Method Returns the Gemini Order Book
	*/

	url := "https://api.gemini.com/v1/book/" + currency

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error Fetching Kraken Order Book")
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error Fetching Kraken Order Book")
	}

	defer res.Body.Close()

	var gb GeminiBook
	json.NewDecoder(res.Body).Decode(&gb)

	bid_kappa, ask_kappa := getGeminiKappa(gb, 20)

	c <- []float64{bid_kappa, ask_kappa}
	w.Done()

}

func getGeminiKappa(gb GeminiBook, depth int) (float64, float64) {

	// Return This
	var bid_kappa float64
	var ask_kappa float64

	for i := 0; i < depth; i++ {

		bid_amount, _ := strconv.ParseFloat(gb.Bids[i].Amount, 64)
		bid_price, _ := strconv.ParseFloat(gb.Bids[i].Price, 64)
		bid_kappa += bid_amount * bid_price

		ask_amount, _ := strconv.ParseFloat(gb.Asks[i].Amount, 64)
		ask_price, _ := strconv.ParseFloat(gb.Asks[i].Price, 64)
		ask_kappa += ask_amount * ask_price

	}

	return bid_kappa, ask_kappa

}
