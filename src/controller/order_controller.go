package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/denizcamalan/PF_FinalProject/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	ordered_count int 
	tot_ordered_amount float64
	order []entities.Order
	months = []entities.Date{}
	total_amount_in_month float64
)
func ListOrders() gin.HandlerFunc {
	return func(c *gin.Context) {	

		session := sessions.Default(c)
		str_order := session.Get("orders")

		if str_order == nil{
			c.JSON(http.StatusOK," NO ORDERED ")
		}else {
			if str_order != nil{
				str_order := session.Get("orders").(string)
				err := json.Unmarshal([]byte(str_order), &order)
				if err != nil{
					log.Println(err,"Unmarshall : ListOrder")
				}
				
				data := map[string]interface{}{
					"Order": order,
				}
				c.JSON(http.StatusOK, data)
			}
		}
	}
}

// buy all of the items into the cart
func BuyCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		if items != nil { 
			var today = entities.Date{Day: time.Now().Day(),Mount: time.Now().Month(),Year: time.Now().Year()}

			ordered_count++
			session := sessions.Default(c)

			var order []entities.Order

			// buy with campaing price if it exist
			if campaign == 0. {
				tot_ordered_amount += tot_card_price_with_VAT
				total_amount_in_month += TotalInMount(tot_card_price_with_VAT,today)
			}else {
				tot_ordered_amount += campaign
				total_amount_in_month += TotalInMount(campaign,today)
			}
			order = append(order, entities.Order{TotalOrder: int64(ordered_count), Orderered_At: entities.Date{
							Day: time.Now().Day(),Mount: time.Now().Month(), Year: time.Now().Year()},
							TotalAmountinMonth: total_amount_in_month, TotalAmount: tot_ordered_amount},)
			bytesOrder, err := json.Marshal(order)
			if err != nil{
				log.Println(err)
			}
			session.Set("orders", string(bytesOrder))

			err3 := session.Save()
			if err3 != nil{log.Println(err3)}
			
		// delete all of the items into the cart after buy items
			for _, id := range items{
				deleteErr := CART.DeleteItem(id.Id)
				tot_card_price = 0
				tot_card_price_with_VAT = 0
				campaign = 0
				if err != nil{
					log.Println(deleteErr,"DeleteAll")
				}
			}
			items = ListItems()
			c.Redirect(http.StatusFound, "/users/orders")
		}else {
			c.JSON(http.StatusBadRequest,"YOUR CART IS EMTY")	
		}
	}
}

// Total ordered amoun in a mount
func TotalInMount(total float64, today entities.Date) float64{
	var total_in_month float64
	months = append(months, today)
	for _,mount := range months{
		if today.Mount == mount.Mount{
			total_in_month = total
		}
	}
	return total_in_month
}