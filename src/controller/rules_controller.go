package controller

import (
	"math"
)

var discount float64

func Campaign_1() float64{

	
	MaxAmount := 1000.

	if ordered % 3 == 0 && totPrice >= MaxAmount{
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

func Campaign_2() float64{

	for _,value := range items {
		if value.Quantity >= 4{
			discount += Discount(value.Price,0.08)
		}else if  value.Quantity < 4{
			discount += 0 
		}
	}
	return discount
}

func Campaign_3() float64{

	MaxInMonth := 9000.0

	if totPrice > MaxInMonth{
		discount = Discount(totVat,0.1)
	}
	return discount
}

func Discount(value, ratio float64) float64{
	value *= ratio
	return value
}

func SelectCampign() float64{

	campaign1 := Campaign_1()
	campaign2 := Campaign_2()
	campaign3 := Campaign_3()
	
	return MaxOfThree(campaign1,campaign2,campaign3)
}

func MaxOfThree(campaign1, campaign2, campaign3 float64) float64{

	result := math.Max(campaign1, campaign2)
	result = math.Max(result, campaign3)
	return result

}