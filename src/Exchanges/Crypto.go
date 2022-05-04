package Exchanges

import "sync"

func GetCryptoOrderBook(currency string, c chan []float64, w *sync.WaitGroup) {

	w.Done()

}
