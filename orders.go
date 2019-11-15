package main

import "encoding/json"

// Orders slice of Order
type Orders []Order

// UnmarshalOrders write Orders from JSON
func UnmarshalOrders(data []byte) (Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal write JSON from Orders
func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Order ...
type Order struct {
	ID            int       `json:"id" db:"id"`
	CourierID     int       `json:"courier_id" db:"courier_id"`
	ClientID      int       `json:"client_id" db:"client_id"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	Consists      []Consist `json:"consists" db:"consists"`
	OrderCost     float64   `json:"order_cost" db:"order_cost"`
	Delivered     bool      `json:"delivered" db:"delivered"`
	DeliveryDelay int       `json:"delivery_delay" db:"delivery_delay"`
	DateStart     string    `json:"date_start" db:"date_start"`
	DateFinish    string    `json:"date_finish" db:"date_finish"`
}

//Consist products of Order (Delivery: 'true' deliver to Client, 'false' return from Client)
type Consist struct {
	Product  string  `json:"product" db:"product"`
	Quantity float64 `json:"quantity" db:"quantity"`
	Delivery bool    `json:"delivery" db:"delivery"`
}
