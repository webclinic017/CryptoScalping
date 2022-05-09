package Orders

import (
	"net/http"
	"time"
)

type FtxClient struct {
	Client *http.Client
	Api    string
	Secret []byte
}

type NewOrderResponse struct {
	Success bool  `json:"success"`
	Result  Order `json:"result"`
}

type OpenOrders struct {
	Success bool    `json:"success"`
	Result  []Order `json:"result"`
}

type Order struct {
	CreatedAt     time.Time `json:"createdAt"`
	FilledSize    float64   `json:"filledSize"`
	Future        string    `json:"future"`
	ID            int64     `json:"id"`
	Market        string    `json:"market"`
	Price         float64   `json:"price"`
	AvgFillPrice  float64   `json:"avgFillPrice"`
	RemainingSize float64   `json:"remainingSize"`
	Side          string    `json:"side"`
	Size          float64   `json:"size"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	ReduceOnly    bool      `json:"reduceOnly"`
	Ioc           bool      `json:"ioc"`
	PostOnly      bool      `json:"postOnly"`
	ClientID      string    `json:"clientId"`
}
