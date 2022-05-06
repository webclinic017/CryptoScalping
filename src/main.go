package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	a "v2/src/Avellaneda"
	e "v2/src/Exchanges"
)

func main() {

	/*
		Each Exchange has a Unique Convention for Currency Pairs
	*/

	coinbase_currency := "ETH-USD"
	// kraken_currency := "ETHUSD"
	gemini_currency := "ETHUSD"
	// crypto_currency := "ETH_USDT"
	ftx_currency := "ETH/USD"

	/*
		Fetch Order Book from each Exchange in GoRoutine
	*/

	num_exchanges := 3
	runtime.GOMAXPROCS(num_exchanges)
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
		Place Order
	*/

	/*
		Order Management
	*/

}
