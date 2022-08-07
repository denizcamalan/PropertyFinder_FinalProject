package entities

import "time"

type Order struct {
	Order_ID       			int64				`json:"_id"`
	Orderered_At   			time.Time			`json:"ordered_on"`
	TotalAmount				float64				`json:"total_price"`
	TotalAmountinMonth     float64				`json:"total_price_in_mounth"`
	Quantity				int64				`json:"quantity"`
}
