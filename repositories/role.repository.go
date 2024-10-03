package repositories


import (
    "app/database"
    "app/models"
    "time"
)

// Busca todos os papéis (roles)
func GetAllRoles() ([]models.Role, error) {
    var roles []models.Role
    _, err := database.DbMap.Select(&roles, "SELECT * FROM roles")
    return roles, err
}
func GetRole(id int64) ([]models.Role,error) {
    var roles []models.Role
    _, err := database.DbMap.Select(&roles, "SELECT * FROM roles WHERE id = $1", id)
    return roles, err
}

// Cria um novo papel (role)
func CreateRole(role models.Role) error {
	role.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
    // Inserção no banco de dados
    err := database.DbMap.Insert(&role)
    return err
}

// Atualiza um papel existente
func UpdateRole(role models.Role, id int64) error {
    // Execução da query de update para atualizar os campos de Role
    println(role.PermissionID)
    _, err := database.DbMap.Exec(`

        UPDATE roles 
        SET name = $1, "permissionId" = $2, updated_at = $3 
        WHERE id = $4`,
        role.Name, role.PermissionID, time.Now().Format("2006-01-02 15:04:05"), id )
    return err
}

// Deleta um papel pelo ID
func DeleteRole(id int64) error {
    
    _, err := database.DbMap.Exec(`

    DELETE FROM roles WHERE id = $1
    `, id)
    return err
}

