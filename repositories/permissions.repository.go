package repositories

import (
    "app/database"
    "app/models"
    "time"
)

// Busca todas as permissões
func GetAllPermissions() ([]models.Permission, error) {
    var permissions []models.Permission
    _, err := database.DbMap.Select(&permissions, "SELECT * FROM permissions")
    return permissions, err
}

func GetPermissions(id int64) ([]models.Permission, error) {
    var permissions []models.Permission
    println(id)
    _, err := database.DbMap.Select(&permissions, "SELECT * FROM permissions WHERE id = $1", id)
    return permissions, err
}


// Cria uma nova permissão
func CreatePermission(permission models.Permission) error {
	permission.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
    // Inserção no banco de dados
    err := database.DbMap.Insert(&permission)
    return err
}

// Atualiza uma permissão existente
func UpdatePermission(permission models.Permission, id int64) error {
    // Execução da query de update para atualizar os campos de Permission
    _, err := database.DbMap.Exec(`
        UPDATE permissions 
        SET name = $1, description = $2, updated_at = $4
        WHERE id = $3`,
        permission.Name, permission.Description, id, time.Now().Format("2006-01-02 15:04:05"))
    return err
}

// Deleta uma permissão pelo ID
func DeletePermission(id int64) error {
    
    _, err := database.DbMap.Exec(`
    
    DELETE FROM permissions WHERE id = $1
    `, id)
    return err
}
