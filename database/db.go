package database

import (
	"app/models"
	"database/sql"
	"log"

	// "fmt"
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

    // Mapeamento das tabelas com as relações
    dbmap.AddTableWithName(models.Customer{}, "customers").SetKeys(true, "ID")
    dbmap.AddTableWithName(models.Role{}, "roles").SetKeys(true, "ID")
    dbmap.AddTableWithName(models.Permission{}, "permissions").SetKeys(true, "ID")
    

    if err = dbmap.CreateTablesIfNotExists(); err != nil {
        log.Fatal("Failed to create tables", err)
    }
    
    // Garantindo que as chaves estrangeiras sejam aplicadas
    _, err = db.Exec(`
BEGIN;


ALTER TABLE IF EXISTS public.customers
    ADD FOREIGN KEY ("roleId")
    REFERENCES public.roles (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.roles
    ADD FOREIGN KEY ("permissionId")
    REFERENCES public.permissions (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
    `)

    if err != nil {
        log.Fatal("Erro ao criar as chaves estrangeiras:", err)
    }

   
	

    DbMap = dbmap
    return dbmap
}

