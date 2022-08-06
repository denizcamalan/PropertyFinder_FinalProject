package controller

//var discount float64

func Campaign_1() float64{
	var discount float64
	items , _ := cartModel.ListAll()
	
	perCarMaxAmount := 1000.
	if i % 3 == 0 && totPrice >= perCarMaxAmount{
		for _,value := range items{
			if value.VAT == 0.08{
				discount = totVat - Discount(totVat,0.1)
			}else if value.VAT == 0.18{
				discount = totVat - Discount(totVat,0.15)
			}else {
				discount = totVat
			}
		}
		return discount
	}else {
		return 0.
	}
}
var alert int
func Campaign_2() float64{
	var discount float64
	items , _ := cartModel.ListAll()

	for _,value := range items {
		if value.Quantity >= 4{
			discount = totVat - Discount(value.Price,0.08)
			alert++
		}
		// else if alert > 0 {
		// 	alert--
		// 	return discount
		// }else {
		// 	discount = 0.
		// }
	}
	return discount
}

func Campaign_3() float64{
	var discount float64

	MaxInMonth := 9000.0

	if totPrice > MaxInMonth{
		discount = totVat - Discount(totVat,0.1)
	}else {
		discount=0
	}
	return discount
}

func Discount(value, ratio float64) float64{
	value *= ratio
	return value
}

func SelectCampign() float64{
	a := Campaign_1()
	b := Campaign_2()
	c := Campaign_3()

	if(a<b && a<c){
        return a
    }else if(b<a && b<c){
        return b
    }else{
        return c
    }
}