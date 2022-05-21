package MonteCarlo

import "math"

func GetStockSimulation(stock_price float64, drift float64, volatility float64, period_length int, simulation_count int) [][]float64 {

	/*
		Input:
		1. Stock Price: Last Recorded Stock Price
		2. Drift: Forcasted Drift Rate
		3. Volatility: Implied Volatility Measure
		4. Period Length: Length of Simulation
		5. Simulation Count: The Total Number of Simulations

		Output:
		1. [][]float64: Stock Simulation List of Lists
	*/

	var BM [][]float64

	rand := GetBoxMullerTransform(period_length, simulation_count)
	dt := 1.0 / float64(period_length)
	count := 0

	c := make(chan []float64, simulation_count)

	for {

		go stockParallel(rand[count], drift, volatility, stock_price, dt, c)
		arr := <-c
		BM = append(BM, arr)
		count++

		if count >= simulation_count {

			return BM

		}

	}

}

func stockParallel(rand []float64, drift float64, volatility float64, stock_price float64, dt float64, c chan []float64) {

	var temp []float64

	for i := 0; i < len(rand); i++ {

		if i == 0 {

			temp = append(temp, stock_price)

		} else {

			temp = append(temp, (temp[i-1] + (drift * temp[i-1] * dt) + (volatility * temp[i-1] * rand[i-1] * math.Sqrt(dt))))

		}

	}

	c <- temp

}
