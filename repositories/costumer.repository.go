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

func GetCustomer(id int64) (models.Customer, error) {
    var customers models.Customer
    _, err := database.DbMap.Select(&customers, "SELECT * FROM customers WHERE id = $1", id)
    return customers, err
}

func CreateCustomer(customer models.Customer) error {
    customer.CreatedAt = time.Now().Format("2006-01-02 15:04:05")    
    err := database.DbMap.Insert(&customer)
    return err
}

func UpdateCustomer(customer models.Customer, id int64) error {
    customer.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
    // Execução da query de update para atualizar os campos do cliente
    _, err := database.DbMap.Exec(`
        UPDATE clients 
        SET fullName = $1, email = $2, zipCode = $3, cpf = $4, birthDate = $5, rg = $6, 
            status = $7, profilePic = $8, appPassword = $9, bond = $10, "roleId" = $11, 
            updated_at = $12
        WHERE id = $13`,
        customer.FullName, customer.Email, customer.ZipCode, customer.CPF, customer.BirthDate, 
        customer.RG, customer.Status, customer.ProfilePic, customer.AppPassword, customer.Bond, 
        customer.RoleID, customer.UpdatedAt, id)
    return err
}

func DeleteCustomer(id int64) error {
    _, err := database.DbMap.Exec("DELETE FROM customers WHERE id = $1", id)
    return err
}
