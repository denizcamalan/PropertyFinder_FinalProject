package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/denizcamalan/PF_FinalProject/models"
	"github.com/denizcamalan/PF_FinalProject/servers"
	"github.com/gin-gonic/gin"
)
var PRODUCT servers.DataBaseServer = &models.ProductModel{}
var productModel models.ProductModel
var products = productModel.ListAll()

func ProductList() gin.HandlerFunc {
	return func(c *gin.Context) {

		products = productModel.ListAll()

		data := map[string]interface{}{
			"products": products,
		}

		c.JSON(http.StatusOK, data)
	}

}

func ProductAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			id, quantity int64
			name, description string
			price, vat float64
		)

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
		
		products = productModel.ListAll()

		for _,value := range products {
			if  intID == value.Id && value.Quantity > 1 {
				err := productModel.UpdateItem(intID,value.Quantity-1)
				if err != nil { log.Println(err, "RemoveProjectItem : UpdateItem")}
				UpdateNewPrice()
				break
			}else {
				PRODUCT.DeleteItem(intID)
				UpdateNewPrice()
			}
		}
		c.Redirect(http.StatusFound, "/users/productlist")
	}
}

