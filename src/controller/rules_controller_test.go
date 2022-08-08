package controller_test

import (
	"testing"

	"github.com/denizcamalan/PF_FinalProject/entities"
)

var (
	quantity int64
	quantity2 int64 
	discount float64
	want_discount float64
	// totVat 		float64
	// tot 		float64
)

func TestCampaign1(t *testing.T){

}

func TestCampaign2(t *testing.T){

	for test := 0; test < 3; test++{

		switch test{
			case 0:
				quantity = 4
				quantity2 = 4
				want_discount = ((1200*0.18)+1200)*0.08 +((1025.99*0.08)+1025.99)*0.08
			continue
			case 1:
				quantity = 3
				quantity2 = 2
				want_discount = 0.
			continue
			case 2:
				quantity = 3
				quantity2 = 4
				want_discount = ((1025.99*0.08)+1025.99)*0.08
			continue
		}
	
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
			Quantity: quantity,
			Description: "phone",
			},
		}

		discount := Campaign_2(items)
		if discount != want_discount {
			t.Errorf("Discounted Value %f, Want Discounted Value %f", discount, want_discount)
		}
	}
}

func TestCampaign3(t *testing.T){

}

func Campaign_1(i int, totPrice, totVat float64, items []entities.Item) float64{

	MaxAmount := 1000.

	if i % 3 == 0 && totPrice >= MaxAmount{
		for _,value := range items{
			if value.VAT == 0.08{
				discount =+ Discount(totVat,0.1)
			}else if value.VAT == 0.18{
				discount =+ Discount(totVat,0.15)
			}else {
				discount += 0
			}
		}
	}
	return discount
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