package entities

// selected item's information into the cart

type Item struct {
	Id				int64			`json:"id"`
	Name 			string			`json:"name"`	
	Price 			float64			`json:"price"`
	VAT 			float64			`json:"vat"`
	Description 	string			`json:"description"`
	Quantity 		int64			`json:"quantity"`
}