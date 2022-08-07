package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/denizcamalan/PF_FinalProject/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func ListOrders() gin.HandlerFunc {
	return func(c *gin.Context) {	

		session := sessions.Default(c)
		str_order := session.Get("orders")

		if str_order == nil{
			c.JSON(http.StatusOK," NO ORDERED ")
		}else {
			ordernum := strconv.Itoa(i)

			if str_order != nil{
				str_order := session.Get("orders").(string)
				var order []entities.Order
				err := json.Unmarshal([]byte(str_order), &order)
				if err != nil{
					log.Println(err,"Unmars ListOrder")
				}
				
				data := map[string]interface{}{
					"Order"+ordernum : order,
					"TotalOrderPrice" : totPrice,
					"TotalOrder" : i,
				}
				c.JSON(http.StatusOK, data)
			}
		}
	}
}

func BuyCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		items = cartModel.ListAll()
		
		session := sessions.Default(c)
		
		var totQua int64
		//var totprice float64
		var order []entities.Order

		if campaign == 0. {
			totPrice += totVat
		}else {
			totPrice = campaign
		}

		order = append(order, entities.Order{Order_ID: 1,Orderered_At: time.Now() , TotalAmount: 100000,TotalAmountinMonth: 3000 , Quantity: totQua})
		bytesOrder, err := json.Marshal(order)
		if err != nil{
			log.Println(err)
		}
		session.Set("orders", string(bytesOrder))

		err3 := session.Save()
		if err3 != nil{log.Println(err3)}
		
		for _, id := range items{
			deleteErr := CART.DeleteItem(id.Id)
			tot = 0
			totVat = 0
			if err != nil{
				log.Println(deleteErr,"DeleteAll")
			}
		}
		i++
		c.Redirect(http.StatusFound, "/users/orders")
	}
}