package services

import (
    "app/models"
    "app/repositories"
)

func GetCustomers() ([]models.Customer, error) {
    return repositories.GetAllCustomers()
}

func GetCustomer(id int64) (models.Customer, error) {
    return repositories.GetCustomer(id)
}

func CreateCustomer(customer models.Customer) error {
    return repositories.CreateCustomer(customer)
}

func UpdateCustomer(customer models.Customer, id int64) error {
    return repositories.UpdateCustomer(customer, id)
}

func DeleteCustomer(id int64) error {
    return repositories.DeleteCustomer(id)
}

