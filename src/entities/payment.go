package entities

import "time"

// order informations

type Order struct {
	TotalOrder      		int64				`json:"total_order"`
	Orderered_At   			Date				`json:"orderered_at"`
	TotalAmount				float64				`json:"total_amount"`
	TotalAmountinMonth     	float64				`json:"total_amount_in_month "`
}

// order date struct
type Date struct{
	Day		int
	Mount 	time.Month
	Year	int
}