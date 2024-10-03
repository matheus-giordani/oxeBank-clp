package models

type Permission struct {
    ID          int64  `db:"id, primarykey, autoincrement" json:"id"`      // Chave primária (inteiro)
    Name        string `db:"name" json:"name"`                            // Nome da permissão (varchar)
    Description string `db:"description" json:"description"`
    CreatedAt  string `db:"created_at" json:"created_at"`
    UpdatedAt  string `db:"updated_at" json:"updated_at"`              // Descrição da permissão (text)
}
