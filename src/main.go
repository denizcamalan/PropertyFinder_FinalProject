package main

import (
	"log"
	"os"

	"github.com/denizcamalan/PF_FinalProject/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)



func main(){

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	store := cookie.NewStore([]byte("mysession"))
  	router.Use(sessions.Sessions("mysession", store))
	routes.UserRoutes(router)

	log.Fatal(router.Run(":" + port))

	
}