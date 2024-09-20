package services

import (
    "app/models"
    "app/repositories"
)

func GetCustomers() ([]models.Customer, error) {
    return repositories.GetAllCustomers()
}

func CreateCustomer(customer models.Customer) error {
    return repositories.CreateCustomer(customer)
}

func UpdateCustomer(customer models.Customer) error {
    return repositories.UpdateCustomer(customer)
}

func DeleteCustomer(id int64) error {
    return repositories.DeleteCustomer(id)
}
