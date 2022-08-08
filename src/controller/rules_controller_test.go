package controller_test

import (
	"testing"

	"github.com/denizcamalan/PF_FinalProject/entities"
)

var (
	//tot_ordered_price float64
	ordered_count int
	discount float64
)

func TestCampaign1(t *testing.T){

	var want_discount float64
	tot_ordered_price := 1500.

	items := []entities.Item{
		{
			Id: 1,
			Name: "Macbook Air",
			Price: 1200,
			VAT: 0.18,
			Description: "computer",
			Quantity: 1,
		},
		{
			Id:	3,
			Name: "Apple 13 Pro",
			Price: 1000,
			VAT: 0.08,
			Quantity: 1,
			Description: "phone",
		},
		{
			Id:	5,
			Name: "Vestel Smart",
			Price: 350,
			VAT: 0.1,
			Quantity: 1,
			Description: "tv",
		},
	}
	for test:= 0; test<4; test++ {

		switch test{
			case 0:
				ordered_count = 4
				want_discount = (1200*1.18*0.15)+(1000*1.08*0.1)
			case 1:
				ordered_count =	6
				want_discount = 0
			case 2:
				ordered_count = 8
				want_discount = (1200*1.18*0.15)+(1000*1.08*0.1)
			case 3:
				ordered_count = 12
				tot_ordered_price =900
				want_discount = 0.	
		}

		discount := Campaign_1(ordered_count,tot_ordered_price,items)
		if discount != want_discount {
			t.Errorf("Discounted Value %f, Want Value %f : Error Case %d", discount, want_discount,test)
		}
	}
}

func Campaign_1(ordered_count int, tot_ordered_price float64, items []entities.Item) float64{

	MaxAmount := 1000.

	if ordered_count % 4 == 0 && tot_ordered_price >= MaxAmount{
		for _,value := range items{
			price_with_VAT := (value.Price*value.VAT)+value.Price
			if value.VAT == 0.08{
				discount += Discount(price_with_VAT,0.1)
			}else if value.VAT == 0.18{
				discount += Discount(price_with_VAT,0.15)
			}else {
				discount += 0
			}
		}
	}else {
		discount = 0
	}
	return discount
}

func TestCampaign2(t *testing.T){
	var want_discount float64
	var quantity int64
	var quantity2 int64 
	
	for test := 0; test<3; test++{
		switch test{
			case 0:
				discount = 0
				quantity = 4
				quantity2 = 4
				want_discount = (1200*1.18+1000*1.08)*0.08
			case 1:
				discount = 0
				quantity = 3
				quantity2 = 2
				want_discount = 0.
			case 2:
				discount = 0
				quantity = 3
				quantity2 = 4
				want_discount = 1000*1.08*0.08
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
				Price: 1000,
				VAT: 0.08,
				Quantity: quantity2,
				Description: "phone",
			},
		}
		discount := Campaign_2(items)
		if discount != want_discount {
			t.Errorf("Discounted Value %f, Want Value %f : Error Case %d", discount, want_discount,test)
		}
	}
}

	func Campaign_2(items []entities.Item) float64{

		for _,value := range items {
			if value.Quantity >= 4{
				discount += Discount((value.Price*value.VAT)+value.Price,0.08)
			}else {
				discount += 0
			}
		}
		return discount
	}	

	func Discount(value, ratio float64) float64{
		value *= ratio
		return value
	}
	
	func TestCampaign3(t *testing.T){

		var want_discount float64
		var tot_card_price_with_VAT = 5000.
		var total_amount_in_mount float64

		for test := 0; test<3; test++{

			switch test{
				case 0:
					total_amount_in_mount = 10000
					want_discount = 500
				case 1:
					total_amount_in_mount = 8000
					want_discount = 0
				case 2:
					total_amount_in_mount = 9000
					want_discount = 0
			}

			orders := entities.Order{
					TotalOrder: 3,
					Orderered_At: entities.Date{},
					TotalAmount: 9000,
					TotalAmountinMonth: total_amount_in_mount,
				}

			discount := Campaign_3(tot_card_price_with_VAT,orders)
			if discount != want_discount {
				t.Errorf("Discounted Value %f, Want Value %f : Error Case %d", discount, want_discount,test)
			}
		}
	}

	func Campaign_3(tot_card_price_with_VAT float64,orders entities.Order) float64{

		MaxInMonth := 9000.0
	
		if orders.TotalAmountinMonth > MaxInMonth{
			discount = Discount(tot_card_price_with_VAT,0.1)
		}else {
			discount = 0
		}
		return discount
	}