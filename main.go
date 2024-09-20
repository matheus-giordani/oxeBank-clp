package main

import (
	"github.com/gin-gonic/gin"
	"app/controllers"
	"app/database"
)

func main() {
	router := gin.Default()


	database.InitDb()


	// Define rotas
    router.GET("/customers", controllers.GetCustomers)
    router.POST("/customers", controllers.CreateCustomer)
    router.PUT("/customers", controllers.UpdateCustomer)
    router.DELETE("/customers/:id", controllers.DeleteCustomer)

    // Starta o servidor
    router.Run(":8080")
}