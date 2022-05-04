package Exchanges

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func GetKrakenOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	/*
		Method Returns the Kraken Order Book
	*/

	url := "https://api.kraken.com/0/public/Depth?pair=XBTUSD"

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

	var kb KrakenBook
	json.NewDecoder(res.Body).Decode(&kb)

	bid_kappa, ask_kappa := getKrakenKappa(kb, 20)

	c <- []float64{bid_kappa, ask_kappa}
	w.Done()

}

func getKrakenKappa(kb KrakenBook, depth int) (float64, float64) {

	// Return This
	var bid_kappa float64
	var ask_kappa float64

	for i := 0; i < depth; i++ {
		bid_kappa += kb.Result.Bids[i][0] * kb.Result.Bids[i][1]
		ask_kappa += kb.Result.Asks[i][0] * kb.Result.Asks[i][1]
	}

	return bid_kappa, ask_kappa

}
