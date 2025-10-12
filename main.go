package main

import (
	"go-eCommerce/initializers"
	"go-eCommerce/middleware"
	"go-eCommerce/controllers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()

	// users
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// products
	r.GET("products", middleware.RequireAuth, controllers.ListProducts)
	r.POST("add-product", middleware.RequireAuth, controllers.AddProduct)
	r.Run()
}