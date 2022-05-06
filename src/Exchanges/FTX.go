package Exchanges

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func GetFTXOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	/*
		Method Returns the FTX Order Book
	*/

	url := "https://ftx.us/api/markets/" + currency + "/orderbook?depth=20"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println("Error Fetching FTX Order Book")
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error Fetching FTX Order Book")
	}

	defer res.Body.Close()

	var fb FTXBook
	json.NewDecoder(res.Body).Decode(&fb)

	best_bid, best_ask, bid_kappa, ask_kappa := getFTXKappa(fb, 20)

	c <- []float64{best_bid, best_ask, bid_kappa, ask_kappa}
	w.Done()

}

func getFTXKappa(fb FTXBook, depth int) (float64, float64, float64, float64) {

	best_bid := fb.Result.Bids[0][0]
	bid_kappa := fb.Result.Bids[0][0] * fb.Result.Bids[0][1]

	best_ask := fb.Result.Asks[0][0]
	ask_kappa := fb.Result.Asks[0][0] * fb.Result.Asks[0][1]

	for i := 1; i < depth; i++ {
		bid_kappa += fb.Result.Bids[i][0] * fb.Result.Bids[i][1]
		ask_kappa += fb.Result.Asks[i][0] * fb.Result.Asks[i][1]
	}

	// fmt.Println("FTX US")
	// fmt.Println("Best Bid: ", best_bid, "Best Ask: ", best_ask)
	// fmt.Println("Bid: ", bid_kappa, "Ask: ", ask_kappa)

	return best_bid, best_ask, bid_kappa, ask_kappa

}
