package Orders

import (
	"encoding/json"
	"log"
)

type OrderTicket struct {
	// Required
	Market string
	Side   string
	Price  float64
	Type   string
	Size   float64

	// Optional
	ReduceOnly bool
	IOC        bool
	PostOnly   bool
}

func (client *FtxClient) PlaceOrder(class *OrderTicket) (NewOrderResponse, error) {

	postBody, _ := json.Marshal(class)
	var newOrderResponse NewOrderResponse

	resp, err := client._post("orders", postBody)

	if err != nil {
		log.Println("Error PlaceOrder", err)
		return newOrderResponse, err
	}
	err = _processResponse(resp, &newOrderResponse)

	return newOrderResponse, err

}

func (client *FtxClient) GetOpenOrders(market string) (OpenOrders, error) {

	var openOrders OpenOrders

	resp, err := client._get("orders?market="+market, []byte(""))

	if err != nil {
		log.Println("Error GetOpenOrders", err)
		return openOrders, err
	}

	err = _processResponse(resp, &openOrders)

	return openOrders, err

}
