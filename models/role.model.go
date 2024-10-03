package models


type Role struct {
    ID           int64  `db:"id, primarykey, autoincrement" json:"id"`      // Chave primária (inteiro)
    Name         string `db:"name" json:"name"`                            // Nome do papel (varchar)
    PermissionID int64  `db:"permissionId" json:"permission_id"`
    CreatedAt  string `db:"created_at" json:"created_at"`
    UpdatedAt  string `db:"updated_at" json:"updated_at"`            // Chave estrangeira para `permissions` (inteiro)
}
