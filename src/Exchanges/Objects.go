package Exchanges

type CoinbaseBook struct {
	Sequence int64       `json:"sequence"`
	Bids     [][]float64 `json:"bids"`
	Asks     [][]float64 `json:"asks"`
}

type KrakenBook struct {
	Error  []string `json:"error"`
	Result struct {
		Data struct {
			Asks [][]float64 `json:"asks"`
			Bids [][]float64 `json:"bids"`
		} `json:"Data"`
	} `json:"result"`
}

type GeminiBook struct {
	Bids []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		Timestamp string `json:"timestamp"`
	} `json:"asks"`
}

type CryptoBook struct {
	Code   int    `json:"code"`
	Method string `json:"method"`
	Result struct {
		Bids [][]float64 `json:"bids"`
		Asks [][]float64 `json:"asks"`
		T    int64       `json:"t"`
	} `json:"result"`
}

type FTXBook struct {
	Success bool `json:"success"`
	Result  struct {
		Asks [][]float64 `json:"asks"`
		Bids [][]float64 `json:"bids"`
	} `json:"result"`
}
