package Exchanges

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func GetCryptoOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	/*
		Method Returns the Crypto Order Book
	*/

	url := "https://uat-api.3ona.co/v2/public/get-book?instrument_name=" + currency + "&depth=20"

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

	var cb CryptoBook
	json.NewDecoder(res.Body).Decode(&cb)

	bid_kappa, ask_kappa := getCryptoKappa(cb, 20)

	c <- []float64{bid_kappa, ask_kappa}
	w.Done()

}

func getCryptoKappa(cb CryptoBook, depth int) (float64, float64) {

	// Return This
	var bid_kappa float64
	var ask_kappa float64

	for i := 0; i < depth; i++ {
		bid_kappa += cb.Result.Bids[i][0] * cb.Result.Bids[i][1]
		ask_kappa += cb.Result.Asks[i][0] * cb.Result.Asks[i][1]
	}

	return bid_kappa, ask_kappa

}
