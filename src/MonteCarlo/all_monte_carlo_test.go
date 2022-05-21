package MonteCarlo

import (
	"fmt"
	"testing"
	"time"
)

func TestBoxMuller(t *testing.T) {

	simulation_length := 100
	simulation_count := 100

	start := time.Now()
	GetBoxMullerTransform(simulation_length, simulation_count)
	end := time.Now()

	fmt.Println("Box Muller Parallel Time Elapsed: ", (end.Sub(start)))

}

func TestHestonVolatility(t *testing.T) {

	implied_vol := 0.20
	vol_floor := 0.10
	vol_reversion := 0.25
	period_length := 100
	simulation_count := 100
	vol_vol := 0.15

	start := time.Now()
	GetHestonVol(implied_vol, vol_floor, vol_reversion, period_length, simulation_count, vol_vol)
	end := time.Now()

	fmt.Println("Heston Volatility Parallel Time Elapsed: ", (end.Sub(start)))

}

func TestOrnsteinUhlenbeck(t *testing.T) {

	reversion_rate := 1.0
	mu := 0.0
	initial := 0.0
	vol := 0.3
	simulation_length := 100
	simulation_count := 100

	start := time.Now()
	GetOrnsteinUhlenback(reversion_rate, mu, initial, vol, simulation_length, simulation_count)
	end := time.Now()

	fmt.Println("Ornstein Uhlenbeck Parallel Time Elapsed: ", (end.Sub(start)))

}

func TestBrownianMotion(t *testing.T) {

	stock_price := 100.0
	drift := 0.01
	vol := 0.25
	period_length := 100
	simulation_count := 100

	start := time.Now()
	GetStockSimulation(stock_price, drift, vol, period_length, simulation_count)
	end := time.Now()

	fmt.Println("Geometric Brownian Motion Parallel Time Elapsed: ", (end.Sub(start)))

}
