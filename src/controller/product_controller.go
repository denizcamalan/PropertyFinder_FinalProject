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
	PRODUCT servers.DataBaseServer = &models.ProductModel{}
	product_result = ListProducts()
)

// list product with JSON format
func ProductList() gin.HandlerFunc {
	return func(c *gin.Context) {

		product_result = ListProducts()

		data := map[string]interface{}{
			"products": product_result,
		}
		c.JSON(http.StatusOK, data)
	}
}
 // add product to market
func ProductAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (id, quantity int64
			name, description string
			price, vat float64)

		// ID
		id1 := c.Request.URL.Query().Get("id")
		intID,err := strconv.ParseInt(id1,10,64)
		if err != nil{log.Println(err, "ProductAdd : strconv : id")}
		id = intID

		// QUANTITY
		quantity1 := c.Request.URL.Query().Get("quantity")
		intID2,err2 := strconv.ParseInt(quantity1,10,64)
		if err2 != nil{log.Println(err2, "ProductAdd : strconv : quantity")}
		quantity = intID2

		// NAME
		name = c.Request.URL.Query().Get("name")

		// DESCRIPTION
		description = c.Request.URL.Query().Get("description")

		// PRICE
		price1 := c.Request.URL.Query().Get("price")
		intID3,err3 := strconv.ParseFloat(price1,64)
		if err3 != nil{log.Println(err3, "ProductAdd : strconv : price")}
		price = intID3

		// VAT
		vat1 := c.Request.URL.Query().Get("vat")
		intID4,err4 := strconv.ParseFloat(vat1,64)
		if err4 != nil{log.Println(err4,"ProductAdd : strconv : vat")}
		vat = intID4

		PRODUCT.AddItem(id,quantity, name, description, price, vat)

		c.Redirect(http.StatusFound, "/users/productlist")
	}

}

// delete product to the market
func RemoveProductItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		productQueryID := c.Request.URL.Query().Get("id")
		if productQueryID == "" {
			log.Println("product id is empty",productQueryID)
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		intID,err := strconv.ParseInt(productQueryID,10,64)
		if err != nil{log.Println(err, "RemoveProjectItem : strconv ")}

		PRODUCT.DeleteItem(intID)
		
		c.Redirect(http.StatusFound, "/users/productlist")
	}
}

// call ListAll function
func ListProducts() []entities.Product{
	products := []entities.Product{}
	for _,value := range PRODUCT.ListAll(){
		products = append(products , value.(entities.Product))
	}
	return products
}