package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/denizcamalan/PF_FinalProject/entities"
	"github.com/denizcamalan/PF_FinalProject/models"
	"github.com/denizcamalan/PF_FinalProject/servers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (	
		a servers.DataBaseServer = &models.CartModel{}
		cartModel models.CartModel
		tot float64
		totVat float64
		campaign float64
		i int
		totPrice float64
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

		userProduct, err := productModel.SelectItem(intID)
		if err != nil {
			log.Println(err, "AddToCard : SelectItem")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		items, err := cartModel.ListAll()
		if err !=nil{
			log.Println(err, "AddToCard : ListAll")
			return
		}
		if items == nil {
			err2 := a.AddItem(userProduct.Id,1,userProduct.Name,userProduct.Description,userProduct.Price,userProduct.VAT )
				if err != nil {
					log.Println(err2,"Primary key")
					return
				}
				UpdateNewPrice()
		}else {
			for _,values := range items {
				if values.Id != intID{
					err2 := a.AddItem(userProduct.Id,1,userProduct.Name,userProduct.Description,
					userProduct.Price,userProduct.VAT )	
					UpdateNewPrice()
					if err != nil {
					log.Println(err2,"Primary key")
					} 
				}else {
					a.UpdateItem(intID,values.Quantity+1)
					UpdateNewPrice()
				}
			}
		}
		c.Redirect(http.StatusFound,"/users/cart")
	}
}

func ListCart()gin.HandlerFunc {
	return func(c *gin.Context) {
		
		items, err := cartModel.ListAll()
		if err !=nil{
			log.Println(err, "ListCard : ListAll")
			c.JSON(http.StatusOK," NO ITEM ")
		}		
		campaign = SelectCampign()

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
		if productQueryID == "" {
			log.Println("product id is empty",productQueryID)
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{
			log.Println(err)
		}
		
		items , _ := cartModel.ListAll()
		log.Println(intID,": intID")
		log.Println("*")
		for _,value := range items {
			if  intID == value.Id && value.Quantity > 1 {
				log.Println(value.Id,": ID : if")
				err := a.UpdateItem(intID,value.Quantity-1)
				if err != nil { log.Println(err, "RemoveToCart : UpdateItem")}
				UpdateNewPrice()
				break
			}else if intID == value.Id && value.Quantity == 1 {
				log.Println(value.Id,": ID : else")
				a.DeleteItem(intID)
				UpdateNewPrice()
			}
		}
		c.Redirect(http.StatusFound, "/users/cart")
	}
}

func ListOrders() gin.HandlerFunc {
	return func(c *gin.Context) {	

		items,_ := cartModel.ListAll()
		if items == nil{
			c.JSON(http.StatusOK," NO ORDERED ")
		}else {
			ordernum := strconv.Itoa(i)
			session := sessions.Default(c)
			str_order := session.Get("orders")

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

		userOrder, err := cartModel.ListAll()
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		session := sessions.Default(c)
		
		var totQua int64
		var totprice float64
		var order []entities.Order

		order = append(order, entities.Order{Order_ID: 1,Orderered_At: time.Now() , Price: totprice, Quantity: totQua})
		bytesOrder, err := json.Marshal(order)
		if err != nil{
			log.Println(err)
		}
		session.Set("orders", string(bytesOrder))

		err3 := session.Save()
		if err3 != nil{
			log.Println(err3)
		}
		if campaign == 0. {
			totPrice += totVat
		}else {
			totPrice = campaign
		}
		
		for _, id := range userOrder{
			deleteErr := a.DeleteItem(id.Id)
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

func Delete(cart []entities.ProductUser, index int64) []entities.ProductUser{
	copy(cart[index:],cart[index+1:])
	return cart[:len(cart)-1]
}

func UpdateNewPrice(){
	item , _ := cartModel.ListAll()
	var totVat1 = 0.
	var tot1 = 0.
	for _,value := range item {
		tot1 += Total(value.Price,value.Quantity)
		totVat1 += TotalWithVAT(value.Price,value.VAT,value.Quantity)
	}
	tot = tot1
	totVat = totVat1
}
