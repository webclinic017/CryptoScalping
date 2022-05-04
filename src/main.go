package main

import (
	"fmt"
	"runtime"
	"sync"
	e "v2/src/Exchanges"
)

func main() {

	/*
		Each Exchange has a Unique Convention for Currency Pairs
	*/

	coinbase_currency := "ETH-USD"
	kraken_currency := "ETHUSD"
	gemini_currency := "ETHUSD"
	crypto_currency := "ETH_USD"
	ftx_currency := "ETH/USD"

	/*
		Fetch Order Book from each Exchange in GoRoutine
	*/

	runtime.GOMAXPROCS(5)

	var coinbase_chan chan []float64
	var kraken_chan chan []float64
	var gemini_chan chan []float64
	var crypto_chan chan []float64
	var ftx_chan chan []float64
	var wg sync.WaitGroup
	wg.Add(5)

	go e.GetCoinbaseOrderBook(coinbase_currency, coinbase_chan, &wg)
	go e.GetKrakenOrderBook(kraken_currency, kraken_chan, &wg)
	go e.GetGeminiOrderBook(gemini_currency, gemini_chan, &wg)
	go e.GetCryptoOrderBook(crypto_currency, crypto_chan, &wg)
	go e.GetFTXOrderBook(ftx_currency, ftx_chan, &wg)
	var count int64

	/*
		Fetch Data from Channels
	*/

	var coinbase_book []float64
	var kraken_book []float64
	var gemini_book []float64
	var crypto_book []float64
	var ftx_book []float64

	go func() {

		for {

			select {
			default:

				if count >= 5 {
					return
				}

			case <-coinbase_chan:

				coinbase_book = <-coinbase_chan

				fmt.Println("Coinbase")
				fmt.Println("Bid: ", coinbase_book[0], "Ask: ", coinbase_book[1])
				count++

			case <-kraken_chan:

				kraken_book = <-kraken_chan
				fmt.Println("Kraken")
				fmt.Println("Bid: ", kraken_book[0], "Ask: ", kraken_book[1])
				count++

			case <-gemini_chan:

				gemini_book = <-gemini_chan

				fmt.Println("Gemini")
				fmt.Println("Bid: ", gemini_book[0], "Ask: ", gemini_book[1])
				count++

			case <-crypto_chan:

				crypto_book = <-crypto_chan

				fmt.Println("Crypto")
				fmt.Println("Bid: ", crypto_book[0], "Ask: ", crypto_book[1])
				count++

			case <-ftx_chan:

				ftx_book = <-ftx_chan

				fmt.Println("FTX US")
				fmt.Println("Bid: ", ftx_book[0], "Ask: ", ftx_book[1])
				count++

			}

		}

	}()

}
