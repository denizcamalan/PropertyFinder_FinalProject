package routes

import (
	"github.com/denizcamalan/PF_FinalProject/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(Router *gin.Engine){

	Router.GET("/users/productlist", controller.ProductList())
	Router.POST("/users/cart/add", controller.AddToCart())
	Router.GET("/users/cart", controller.ListCart())
	Router.POST("/users/cart/remove", controller.RemoveToCart())
	Router.POST("/users/cart/buy", controller.BuyCart())
	Router.GET("/users/orders", controller.ListOrders())
	Router.POST("/users/product/add",controller.ProductAdd())
	Router.POST("/users/product/delete",controller.RemoveProductItem())
}