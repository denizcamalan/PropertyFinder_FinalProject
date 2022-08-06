package entities

type Product struct{
	Id				int64			`json:"id"`
	Name 			string			`json:"name"`	
	Price 			float64			`json:"price"`
	VAT 			float64			`json:"vat"`
	Quantity 		int64			`json:"quantity"`
	Description 	string			`json:"description"`
}