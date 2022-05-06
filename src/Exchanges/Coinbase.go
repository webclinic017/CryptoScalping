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

	best_bid, best_ask, bid_kappa, ask_kappa := getCoinbaseKappa(cb)

	c <- []float64{best_bid, best_ask, bid_kappa, ask_kappa}
	w.Done()

}

func getCoinbaseKappa(cb CoinbaseBook) (float64, float64, float64, float64) {

	best_bid := cb.Bids[0][0]
	bid_kappa := cb.Bids[0][0] * cb.Bids[0][1]

	best_ask := cb.Asks[0][0]
	ask_kappa := cb.Asks[0][0] * cb.Asks[0][1]

	for i := 1; i < len(cb.Bids); i++ {
		bid_kappa += cb.Bids[i][0] * cb.Bids[i][1]
		ask_kappa += cb.Asks[i][0] * cb.Asks[i][1]
	}

	// fmt.Println("Coinbase")
	// fmt.Println("Best Bid: ", best_bid, "Best Ask: ", best_ask)
	// fmt.Println("Bid: ", bid_kappa, "Ask: ", ask_kappa)

	return best_bid, best_ask, bid_kappa, ask_kappa

}
