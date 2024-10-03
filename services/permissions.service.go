package services

import (
    "app/models"
    "app/repositories"
)



// Serviço para buscar todos os papéis (roles)
func GetAllPermissions() ([]models.Permission, error) {
	return repositories.GetAllPermissions()    
    
}

func GetPermissions(id int64) ([]models.Permission, error){
	return repositories.GetPermissions(id)
}


// Serviço para criar um novo papel (Permissions)
func CreatePermissions(Permissions models.Permission) error {
    // Validação adicional ou lógica de negócios pode ser adicionada aqui
    err := repositories.CreatePermission(Permissions)
    return err
}

// Serviço para atualizar um papel (Permissions)
func UpdatePermissions(permission models.Permission, id int64) error {
    // Validação adicional ou lógica de negócios pode ser feita aqui antes de atualizar
    err := repositories.UpdatePermission(permission, id)
    return err
}

// Serviço para deletar um papel (role)
func DeletePermission(id int64) error {
    // Você pode adicionar verificação de dependências ou outras lógicas de negócios antes de deletar
    err := repositories.DeletePermission(id)
    return err
}