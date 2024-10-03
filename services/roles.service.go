package services

import (
    "app/models"
    "app/repositories"
)


// Serviço para buscar todos os papéis (roles)
func GetAllRoles() ([]models.Role, error) {
	return repositories.GetAllRoles()    
    
}

func GetRole(id int64) ([]models.Role, error){
	return repositories.GetRole(id)
}


// Serviço para criar um novo papel (role)
func CreateRole(role models.Role) error {
    // Validação adicional ou lógica de negócios pode ser adicionada aqui
    err := repositories.CreateRole(role)
    return err
}

// Serviço para atualizar um papel (role)
func UpdateRole(role models.Role, id int64) error {
    // Validação adicional ou lógica de negócios pode ser feita aqui antes de atualizar
    err := repositories.UpdateRole(role, id)
    return err
}

// Serviço para deletar um papel (role)
func DeleteRole(id int64) error {
	
    // Você pode adicionar verificação de dependências ou outras lógicas de negócios antes de deletar
    err := repositories.DeleteRole(id)
    return err
}