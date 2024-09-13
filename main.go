package main

import (
	"go-gin-jwt/controllers"
	"go-gin-jwt/intializers"
	"go-gin-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	//load env
	initializers.LoadEnv()
	//connect database
    initializers.ConnectDatabase()

	//init router
	r := gin.Default()

  	r.POST("/register", controllers.Register) // Request register
	r.POST("/login", controllers.Login) // Request login
	r.GET("/validasi",middleware.Requireauth, controllers.Validasi) // Request validasi login

  	r.Run()
}