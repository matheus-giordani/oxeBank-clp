package models

type Customer struct {
    ID         int64  `db:"id, primarykey, autoincrement" json:"id"`
    Name       string `db:"name" json:"name"`
    Email      string `db:"email" json:"email"`
    Phone      string `db:"phone" json:"phone"`
    CreatedAt  string `db:"created_at" json:"created_at"`
    UpdatedAt  string `db:"updated_at" json:"updated_at"`
}