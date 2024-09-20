package database


import (
	"database/sql"
	"log"
	// "fmt"
	"app/models"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"	

)


var DbMap *gorp.DbMap

func InitDb() *gorp.DbMap {
    db, err := sql.Open("postgres", "user=user dbname=db password=123 host=postgresql sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to the database", err)
    }
	 
	// Verifica a conexão
    if err = db.Ping(); err != nil {
        log.Fatal("Erro ao verificar a conexão com o banco:", err)
    }

    // Se a conexão for bem-sucedida, imprime uma mensagem
    log.Println("Conexão com o banco de dados estabelecida com sucesso!")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.AddTableWithName(models.Customer{}, "customers").SetKeys(true, "ID")

    if err = dbmap.CreateTablesIfNotExists(); err != nil {
        log.Fatal("Failed to create tables", err)
    }
	

    DbMap = dbmap
    return dbmap
}

