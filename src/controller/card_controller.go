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
		tot_card_price float64
		tot_card_price_with_VAT float64  
		campaign float64
		items = ListItems()
		CART servers.DataBaseServer = &models.CartModel{}
	)

// Add item to the cart
func AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Request.URL.Query().Get("id")
		if productQueryID == ""{
			log.Println("product id is empty",productQueryID)
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		// convert strng id to int
		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{
			CART.DeleteItem(intID)
			log.Println(err)
		}
		// Select items to products table 
		userCart := PRODUCT.SelectItem(intID).(entities.Product)

		// list selected items and make slice
		items = ListItems()

		// if cart empty add item
		if items == nil {
			err2 := CART.AddItem(userCart.Id,1,userCart.Name,userCart.Description,userCart.Price,userCart.VAT )
				if err != nil {
					log.Println(err2,"Primary key")
					return
				}
				// update price or campaign
				UpdateNewPrice()
				FindCampgainPrice()
		}else {
			for _,values := range items {
				if values.Id != intID{
					err2 := CART.AddItem(userCart.Id,1,userCart.Name,userCart.Description,
						userCart.Price,userCart.VAT )
					// update price or campaign	
					UpdateNewPrice()
					FindCampgainPrice()
					if err != nil {
					log.Println(err2,"Primary key")
					} 
				}else {
					// if item exist in the cart ++quantity
					CART.UpdateItem(intID,values.Quantity+1)
					// update price or campaign
					UpdateNewPrice()
					FindCampgainPrice()
				}
			}
		}
		// directed to cart
		c.Redirect(http.StatusFound,"/users/cart")
	}
}

// list cart
func ListCart()gin.HandlerFunc {
	return func(c *gin.Context) {
		
		items = ListItems()

		data := map[string]interface{}{
			"Cart": items,
			"TotalPrice": tot_card_price,
			"TotalwithVAT": tot_card_price_with_VAT,
			"_CampaignPrice": campaign,
		}
		// show data JSON format
		c.JSON(http.StatusOK, data)
	}
}

// delete item from cart
func RemoveToCart()gin.HandlerFunc {
	return func(c *gin.Context) {

		productQueryID := c.Request.URL.Query().Get("id")

		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{
			log.Println(err)
		}
		
		items = ListItems()

		for _,value := range items {
			// if item already exist change quantity
			if  intID == value.Id && value.Quantity > 1 {
				err := CART.UpdateItem(intID,value.Quantity-1)
				if err != nil { log.Println(err, "RemoveToCart : UpdateItem")}
				// update price or campaign for a new quantity
				UpdateNewPrice()
				FindCampgainPrice()
				break
			}else if intID == value.Id && value.Quantity == 1 {
				CART.DeleteItem(intID)
				// update price or campaign for a new quantity
				UpdateNewPrice()
				FindCampgainPrice()
			}
		}
		// directed to listAll
		c.Redirect(http.StatusFound, "/users/cart")
	}
}

// total cart amount wit VAT
func TotalWithVAT(price, VAT float64, quantity int64) float64 {
	result := 0.
		withVAT := price * VAT
		result += (price + withVAT) * float64(quantity) 
	return result
}

// total cart amounT without VAT
func Total(price float64, quantity int64) float64 {
	result := 0.
		result = price * float64(quantity) 
	return result
}

// update total card price
func UpdateNewPrice(){
	var tot_card_price_with_VAT1 = 0.
	var tot_card_price1 = 0.

	items = ListItems()

	for _,value := range items {
		tot_card_price1 += Total(value.Price,value.Quantity)
		tot_card_price_with_VAT1 += TotalWithVAT(value.Price,value.VAT,value.Quantity)
	}
	tot_card_price = tot_card_price1
	tot_card_price_with_VAT = tot_card_price_with_VAT1
}

// find to campaign is exist or not
func FindCampgainPrice(){
	campaign = tot_card_price_with_VAT - SelectCampign()
	if campaign == tot_card_price_with_VAT || SelectCampign() == 0 {
		campaign = 0
	}
}

// call ListAll function
func ListItems() []entities.Item{
	var item []entities.Item
	for _,value := range CART.ListAll(){
		item = append(item , value.(entities.Item))
	}
	return item
}