package MonteCarlo

import "math"

/*
	Input:
	1. Reversion Rate: The Speed of Mean Reversion
	2. Mu: The Average
	3. Initial: Initial Value, Price or Vol
	4. Volatility: Forecasted Volatilty (Try Heston)
	5. Simulation Length: The Length of the Simulation
	6. Simulation Count: The Total Number of Simulations

	Output:
	1. [][]float64: Mean Reverting Monte Carlo
*/

func GetOrnsteinUhlenback(reversion_rate float64, mu float64, initial float64, volatility float64, simulation_length int, simulation_count int) [][]float64 {

	var simulation [][]float64

	rand := GetBoxMullerTransform(simulation_length, simulation_count)
	dt := 1.0 / float64(simulation_length)
	count := 0

	c := make(chan []float64, simulation_count)

	for {

		go ornsteinParallel(rand[count], reversion_rate, mu, initial, volatility, dt, simulation_length, c)
		arr := <-c
		simulation = append(simulation, arr)
		count++

		if count >= simulation_count {

			return simulation

		}

	}

}

func ornsteinParallel(rand []float64, reversion_rate float64, mu float64, initial float64, volatility float64, dt float64, simulation_length int, c chan []float64) {

	var temp []float64

	for i := 0; i < simulation_length; i++ {

		if i == 0 {

			temp = append(temp, initial)

		} else {

			temp = append(temp, (temp[i-1] + (reversion_rate * (mu - temp[i-1]) * dt) + (volatility * math.Sqrt(dt) * rand[i-1])))

		}

	}

	c <- temp

}
