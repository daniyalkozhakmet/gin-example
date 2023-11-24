package main

import (
	"fmt"
	"gin-example/controllers"
	"gin-example/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.POST("/ping", controllers.SignUp)
	router.Run()
	fmt.Println("Server running")
}