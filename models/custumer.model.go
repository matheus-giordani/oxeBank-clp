package models

type Customer struct {
	ID          int64  `db:"id, primarykey, autoincrement" json:"id"` // Chave primária (inteiro)
	FullName    string `db:"fullName" json:"full_name"`               // Nome completo (varchar)
	Email       string `db:"email" json:"email"`                      // Email (varchar)
	ZipCode     string `db:"zipCode" json:"zip_code"`                 // CEP (varchar)
	CPF         int64  `db:"cpf" json:"cpf"`                          // CPF (integer)
	BirthDate   string `db:"birthDate" json:"birth_date"`             // Data de nascimento (date)
	RG          string `db:"rg" json:"rg"`                            // RG (varchar)
	Status      string `db:"status" json:"status"`                    // Status (enum)
	ProfilePic  string `db:"profilePic" json:"profile_pic"`           // Foto de perfil (text)
	AppPassword string `db:"appPassword" json:"app_password"`         // Senha do app (text)
	Bond        string `db:"bond" json:"bond"`                        // Vínculo (enum)
	RoleID      int64  `db:"roleId" json:"role_id"`                   // Chave estrangeira para a tabela `roles` (inteiro)
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
}
