package repositories

import (
    "app/database"
    "app/models"
    "time"
)

func GetAllCustomers() ([]models.Customer, error) {
    var customers []models.Customer
    _, err := database.DbMap.Select(&customers, "SELECT * FROM customers")
    return customers, err
}

func CreateCustomer(customer models.Customer) error {
    customer.CreatedAt = time.Now().Format("2006-01-02 15:04:05")    
    err := database.DbMap.Insert(&customer)
    return err
}

func UpdateCustomer(customer models.Customer) error {
    customer.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
    _, err := database.DbMap.Update(&customer)
    return err
}

func DeleteCustomer(id int64) error {
    _, err := database.DbMap.Exec("DELETE FROM customers WHERE id = $1", id)
    return err
}
