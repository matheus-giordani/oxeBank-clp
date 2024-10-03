package controllers

import (
    "net/http"
    "strconv"
    "app/models"
    "app/services"
    "github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
    customers, err := services.GetAllRoles()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customers)
}


func GetRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var customer []models.Role
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
	if customer, err = services.GetRole(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customer)
}



func CreateRole(c *gin.Context) {
    var customer models.Role
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.CreateRole(customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, customer)
}

func UpdateRole(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var customer models.Role
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.UpdateRole(customer,id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customer)
}

func DeleteRole(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := services.DeleteRole(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
