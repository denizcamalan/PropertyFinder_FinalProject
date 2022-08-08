package entities

import "time"

type Order struct {
	TotalOrder      		int64				`json:"total_order"`
	Orderered_At   			Date				`json:"orderered_at"`
	TotalAmount				float64				`json:"total_amount"`
	TotalAmountinMonth     	float64				`json:"total_amount_in_month "`
}

type Date struct{
	Day		int
	Mount 	time.Month
	Year	int
}