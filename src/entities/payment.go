package entities

import "time"

type Order struct {
	Order_ID       	int64				`json:"_id"`
	Orderered_At   	time.Time			`json:"ordered_on"`
	Price          	float64				`json:"total_price"`
	Quantity		int64				`json:"quantity"`
}

type Entities struct{
	Product			Product 
	ProductUser 	ProductUser
	Order 			Order
}