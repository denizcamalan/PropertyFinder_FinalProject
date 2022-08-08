package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/denizcamalan/PF_FinalProject/entities"
	"github.com/denizcamalan/PF_FinalProject/models"
	"github.com/denizcamalan/PF_FinalProject/servers"
	"github.com/gin-gonic/gin"
)

var (	
		tot float64
		totVat float64
		campaign float64
		cartModel models.CartModel
		items = cartModel.ListAll()
		CART servers.DataBaseServer = &models.CartModel{}
	)

func AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Request.URL.Query().Get("id")
		if productQueryID == "" {
			log.Println("product id is empty",productQueryID)
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{
			log.Println(err)
		}

		userProduct, err2 := productModel.SelectItem(intID)
		if err2 != nil {
			log.Println(err2, "AddToCard : SelectItem")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		items = cartModel.ListAll()
		if items == nil {
			err2 := CART.AddItem(userProduct.Id,1,userProduct.Name,userProduct.Description,userProduct.Price,userProduct.VAT )
				if err != nil {
					log.Println(err2,"Primary key")
					return
				}
				UpdateNewPrice()
				campaign = totVat - SelectCampign()
		}else {
			for _,values := range items {
				if values.Id != intID{
					err2 := CART.AddItem(userProduct.Id,1,userProduct.Name,userProduct.Description,
					userProduct.Price,userProduct.VAT )	
					UpdateNewPrice()
					campaign = totVat - SelectCampign()
					if err != nil {
					log.Println(err2,"Primary key")
					} 
				}else {
					CART.UpdateItem(intID,values.Quantity+1)
					UpdateNewPrice()
					campaign = totVat - SelectCampign()
				}
			}
		}
		c.Redirect(http.StatusFound,"/users/cart")
	}
}

func ListCart()gin.HandlerFunc {
	return func(c *gin.Context) {
		
		items = cartModel.ListAll()

		if campaign == totVat {
			campaign = 0
		}

		data := map[string]interface{}{
			"Cart": items,
			"TotalPrice": tot,
			"TotalwithVAT": totVat,
			"_CampaignPrice": campaign,
		}

		c.JSON(http.StatusOK, data)
	}
}

func RemoveToCart()gin.HandlerFunc {
	return func(c *gin.Context) {

		productQueryID := c.Request.URL.Query().Get("id")

		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{
			log.Println(err)
		}
		
		log.Println(intID,": intID")
		log.Println("*")
		items = cartModel.ListAll()
		for _,value := range items {
			if  intID == value.Id && value.Quantity > 1 {
				log.Println(value.Id,": ID : if")
				err := CART.UpdateItem(intID,value.Quantity-1)
				if err != nil { log.Println(err, "RemoveToCart : UpdateItem")}
				UpdateNewPrice()
				break
			}else if intID == value.Id && value.Quantity == 1 {
				log.Println(value.Id,": ID : else")
				CART.DeleteItem(intID)
				UpdateNewPrice()
			}
		}
		c.Redirect(http.StatusFound, "/users/cart")
	}
}

func TotalWithVAT(price, VAT float64, quantity int64) float64 {
	result := 0.
		withVAT := price * VAT
		result += (price + withVAT) * float64(quantity) 
	return result
}

func Total(price float64, quantity int64) float64 {
	result := 0.
		result = price * float64(quantity) 
	return result
}

func Delete(cart []entities.Item, index int64) []entities.Item{
	copy(cart[index:],cart[index+1:])
	return cart[:len(cart)-1]
}

func UpdateNewPrice(){
	var totVat1 = 0.
	var tot1 = 0.
	items = cartModel.ListAll()
	for _,value := range items {
		tot1 += Total(value.Price,value.Quantity)
		totVat1 += TotalWithVAT(value.Price,value.VAT,value.Quantity)
	}
	tot = tot1
	totVat = totVat1
}
