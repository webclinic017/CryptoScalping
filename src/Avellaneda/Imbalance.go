package Avellaneda

func OrderBookImbalance(bid_liquidity float64, best_bid float64, ask_liquidity float64, best_ask float64) float64 {

	/*
	   Methods Returns the Weighted Midpoint

	   Input:
	   1. Bid Kappa
	   2. Best Bid
	   3. Ask Kappa
	   4. Best Ask

	   Output:
	   1. The Weighted Midpoint
	*/

	imbalance := bid_liquidity / (bid_liquidity + ask_liquidity)
	weighted_midpoint := (imbalance * best_ask) + ((1 - imbalance) * best_bid)

	return weighted_midpoint

}

func OrderBookSkew(order_books [][]float64) bool {

	var isSkewed []bool

	for i := 0; i < len(order_books); i++ {

		if order_books[i][0] < order_books[i][1] {
			isSkewed = append(isSkewed, true)
		} else {
			isSkewed = append(isSkewed, false)
		}

	}

	skew := true

	for i := 0; i < len(isSkewed); i++ {

		skew = skew && isSkewed[i]

	}

	return skew

}
