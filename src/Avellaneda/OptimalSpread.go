package Avellaneda

import "math"

func GetReservationPrice(midpoint float64, inventory_target float64, gamma float64, sigma float64, tau float64) float64 {

	/*

	 */

	reservation_price := midpoint - (inventory_target * gamma * math.Pow(sigma, 2) * tau)

	return reservation_price

}

func GetOptimalSpread(midpoint float64, gamma float64, kappa float64, sigma float64, tau float64) float64 {

	/*

	 */

	optimal_spread := (gamma * math.Pow(sigma, 2) * tau) + ((2.0 / gamma) * math.Log(1.0+(gamma/kappa)))

	return optimal_spread

}
