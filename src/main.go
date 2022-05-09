package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	a "v2/src/Avellaneda"
	e "v2/src/Exchanges"
)

/*
	Global Variables:
	Each Exchange has a Unique Convention for Currency Pairs

	Trade Parameters Initialized to False
*/

var coinbase_currency string = "ETH-USD"
var kraken_currency string = "ETHUSD"
var gemini_currency string = "ETHUSD"
var crypto_currency string = "ETH_USDT"
var ftx_currency string = "ETH/USD"

var isLong bool = false
var trade_size float64 = 1.0

func main() {

	fmt.Println("Crypto Scalper Starting")

	// Input Api Key
	// var api_key string
	// fmt.Println("Please Enter Api Key: ")
	// fmt.Scanln(&api_key)

	// // Input Api Secret
	// var api_secret string
	// fmt.Println("Please Enter Api Secret: ")
	// fmt.Scanln(&api_secret)

	// // Initialize Client
	// client := o.New(api_key, api_secret)

	// Set Max Threads
	num_exchanges := 3
	runtime.GOMAXPROCS(num_exchanges)

	// Create Ticker
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {

		/*
			Fetch Order Book from each Exchange in GoRoutine
		*/

		start := time.Now()

		coinbase_chan := make(chan []float64, 1)
		// kraken_chan := make(chan []float64, 1)
		gemini_chan := make(chan []float64, 1)
		// crypto_chan := make(chan []float64, 1)
		ftx_chan := make(chan []float64, 1)

		var wg sync.WaitGroup
		wg.Add(num_exchanges)

		go e.GetCoinbaseOrderBook(coinbase_currency, coinbase_chan, &wg)
		// go e.GetKrakenOrderBook(kraken_currency, kraken_chan, &wg)
		go e.GetGeminiOrderBook(gemini_currency, gemini_chan, &wg)
		// go e.GetCryptoOrderBook(crypto_currency, crypto_chan, &wg)
		go e.GetFTXOrderBook(ftx_currency, ftx_chan, &wg)

		wg.Wait()
		end := time.Now()

		fmt.Println("Order Book Routines: ", end.Sub(start))
		fmt.Println("")

		/*
			Fetch Data from Channels
		*/

		var order_books [][]float64

		coinbase_book := <-coinbase_chan
		fmt.Println("Coinbase")
		fmt.Println("Best Bid: ", coinbase_book[0], "Best Ask: ", coinbase_book[1])
		fmt.Println("Bid: ", coinbase_book[2], "Ask: ", coinbase_book[3])
		coinbase_midpoint := (coinbase_book[0] + coinbase_book[1]) / 2.0
		fmt.Println("Midpoint: ", coinbase_midpoint)
		coinbase_weighted_midpoint := a.OrderBookImbalance(coinbase_book[2], coinbase_book[0], coinbase_book[3], coinbase_book[1])
		fmt.Println("Weighted Midpoint: ", coinbase_weighted_midpoint)
		order_books = append(order_books, []float64{coinbase_midpoint, coinbase_weighted_midpoint})
		fmt.Println("")

		// 	kraken_book := <-kraken_chan
		// 	fmt.Println("Kraken")
		// 	fmt.Println("Bid: ", kraken_book[0], "Ask: ", kraken_book[1])
		// fmt.Println("")

		gemini_book := <-gemini_chan
		fmt.Println("Gemini")
		fmt.Println("Best Bid: ", gemini_book[0], "Best Ask: ", gemini_book[1])
		fmt.Println("Bid: ", gemini_book[2], "Ask: ", gemini_book[3])
		gemini_midpoint := (gemini_book[0] + gemini_book[1]) / 2.0
		fmt.Println("Midpoint: ", gemini_midpoint)
		gemini_weighted_midpoint := a.OrderBookImbalance(gemini_book[2], gemini_book[0], gemini_book[3], gemini_book[1])
		fmt.Println("Weighted Midpoint: ", gemini_weighted_midpoint)
		order_books = append(order_books, []float64{gemini_midpoint, gemini_weighted_midpoint})
		fmt.Println("")

		// 	crypto_book := <-crypto_chan
		// 	fmt.Println("Crypto")
		// 	fmt.Println("Bid: ", crypto_book[0], "Ask: ", crypto_book[1])
		// fmt.Println("")

		ftx_book := <-ftx_chan
		fmt.Println("FTX US")
		fmt.Println("Best Bid: ", ftx_book[0], "Best Ask: ", ftx_book[1])
		fmt.Println("Bid: ", ftx_book[2], "Ask: ", ftx_book[3])
		ftx_midpoint := (ftx_book[0] + ftx_book[1]) / 2.0
		fmt.Println("Midpoint: ", ftx_midpoint)
		ftx_weighted_midpoint := a.OrderBookImbalance(ftx_book[2], ftx_book[0], ftx_book[3], ftx_book[1])
		fmt.Println("Weighted Midpoint: ", ftx_weighted_midpoint)
		order_books = append(order_books, []float64{ftx_midpoint, ftx_weighted_midpoint})
		fmt.Println("")

		/*
			Check for Order Book Skew
		*/

		isSkewed := a.OrderBookSkew(order_books)
		fmt.Println("Order Book Skew: ", isSkewed)
		fmt.Println("")

		/*
			Enter Long Position
		*/

		// var OT o.OrderTicket

		// // if isSkewed && !isLong {

		// // 	// Set Variables
		// // 	OT.Market = ftx_currency
		// // 	OT.Side = "buy"
		// // 	OT.Price = ftx_midpoint
		// // 	OT.Type = "limit"
		// // 	OT.Size = trade_size

		// // 	resp, err := client.PlaceOrder(&OT)

		// // 	if err != nil {
		// // 		log.Println(err)
		// // 	}

		// // 	fmt.Println("Order Result: ", resp.Success)

		// // 	go func() {

		// // 		resp, err := client.GetOpenOrders(ftx_currency)

		// // 		if err != nil {
		// // 			log.Println(err)
		// // 		}

		// // 		fmt.Println("Open Orders: ", resp.Success)

		// // 	}()

		// // }

		// /*
		// 	Order Management
		// */

		// if isLong {

		// }

	}

}
