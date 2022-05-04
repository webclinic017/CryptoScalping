package Exchanges

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func GetCoinbaseOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	/*
		Method Returns the Coinbase Order Book
	*/

	url := "https://api.exchange.coinbase.com/products/" + currency + "/book?level=2"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error Fetching Coinbase Order Book")
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error Fetching Coinbase Order Book")
	}

	defer res.Body.Close()

	var cb CoinbaseBook
	json.NewDecoder(res.Body).Decode(&cb)

	bid_kappa, ask_kappa := getCoinbaseKappa(cb)

	c <- []float64{bid_kappa, ask_kappa}
	w.Done()

}

func getCoinbaseKappa(cb CoinbaseBook) (float64, float64) {

	// Return This
	var bid_kappa float64
	var ask_kappa float64

	var bid []float64
	var bid_size []float64

	for i := 0; i < len(cb.Bids); i++ {
		bid = append(bid, cb.Bids[i][0])
		bid_size = append(bid_size, cb.Bids[i][1])
	}

	var ask []float64
	var ask_size []float64

	for i := 0; i < len(cb.Asks); i++ {
		ask = append(ask, cb.Asks[i][0])
		ask_size = append(ask_size, cb.Asks[i][1])
	}

	return bid_kappa, ask_kappa

}
