package controller_test

import (
	"testing"

	"github.com/denizcamalan/PF_FinalProject/entities"
)

var (
	discount 	float64
	totVat 		float64
	tot 		float64
)

func TestCampaign2(t *testing.T){
	var quantity int64 = 0
	var quantity2 int64 = 0
	items := []entities.Item{
		{
		Id: 1,
		Name: "Macbook Air",
		Price: 1200,
		VAT: 0.18,
		Description: "computer",
		Quantity: quantity,
		},
		{
		Id:	3,
		Name: "Apple 13 Pro",
		Price: 1025.99,
		VAT: 0.08,
		Quantity: quantity2,
		Description: "phone",
		}}

	discount = Campaign_2(items)
	want_discount := 201.925536

	if discount != want_discount && quantity == 4 && quantity2 == 4{
        t.Errorf("Discounted Value %f, Want Discounted Value %f", discount, want_discount)
    }

	discount = Campaign_2(items)
	want_discount = 0.

	if discount != want_discount && quantity == 3 && quantity2 == 3 {
        t.Errorf("Discounted Value %f, Want Discounted Value %f", discount, want_discount)
    }

	discount = Campaign_2(items)
	want_discount = 0.
	quantity = 3
	quantity2 = 4
	if discount != want_discount {
        t.Errorf("Discounted Value %f, Want Discounted Value %f", discount, want_discount)
    }
}


	func Campaign_2(items []entities.Item) float64{

			
		for _,value := range items {
			if value.Quantity >= 4{
				price_with_VAT := (value.Price*value.VAT)+value.Price
				discount += Discount(price_with_VAT,0.08)
			}else if  value.Quantity < 4{
				discount += 0 
			}
		}
		return discount
	}	

	func Discount(value, ratio float64) float64{
		value *= ratio
		return value
	}
	
	// func SelectCampign() float64{
	
	// 	campaign1 := Campaign_1()
	// 	campaign2 := Campaign_2()
	// 	campaign3 := Campaign_3()
		
	// 	return MaxOfThree(campaign1,campaign2,campaign3)
	// }
	
	// func MaxOfThree(campaign1, campaign2, campaign3 float64) float64{
	
	// 	result := math.Max(campaign1, campaign2)
	// 	result = math.Max(result, campaign3)
	// 	return result
	
	// }

		// for _,value := range items {
	// 	tot += controller.Total(value.Price,value.Quantity)
	// 	totVat += controller.TotalWithVAT(value.Price,value.VAT,value.Quantity)
	// }