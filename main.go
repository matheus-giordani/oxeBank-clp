package main

import (
	"github.com/gin-gonic/gin"
	"app/controllers"
	"app/database"
)

func main() {
	router := gin.Default()


	database.InitDb()


	// Define customers
    router.GET("/customers", controllers.GetCustomers)
    router.POST("/customers", controllers.CreateCustomer)
    router.PUT("/customers/:id", controllers.UpdateCustomer)
    router.DELETE("/customers/:id", controllers.DeleteCustomer)

	// Rotas para Role
	router.GET("/roles", controllers.GetRoles)               // Obter todos os papéis
	router.GET("/roles/:id", controllers.GetRole)       // Obter um papel por ID
	router.POST("/roles", controllers.CreateRole)            // Criar um novo papel
	router.PUT("/roles/:id", controllers.UpdateRole)        // Atualizar um papel por ID
	router.DELETE("/roles/:id", controllers.DeleteRole)     // Deletar um papel por ID

	// Rotas para Permission
	router.GET("/permissions", controllers.GetPermissions)   // Obter todas as permissões
	router.GET("/permissions/:id", controllers.GetPermission) // Obter uma permissão por ID
	router.POST("/permissions", controllers.CreatePermission) // Criar uma nova permissão
	router.PUT("/permissions/:id", controllers.UpdatePermission) // Atualizar uma permissão por ID
	router.DELETE("/permissions/:id", controllers.DeletePermission) // Deletar uma permissão por ID


    // Starta o servidor
    router.Run(":8081")
}