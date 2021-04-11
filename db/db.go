package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const CONNECTION_STRING = "user=docker password=docker dbname=neoway host=db sslmode=disable"

type Store struct {
	CPF                string
	Privado            bool
	Incompleto         bool
	DataUltimaCompra   time.Time
	TicketMedio        float64
	TicketUltimaCompra float64
	LojaMaisFrequente  string
	LojaUltimaCompra   string
	Invalido           bool
}

//Creates Store table
func createTable(db *sql.DB) {
	create := `CREATE TABLE IF NOT EXISTS "store"  (
		"id" SERIAL PRIMARY KEY,
		"cpf" varchar(20),
		"privado" boolean,
		"incompleto" boolean,
		"data_ultima_compra" date,
		"ticket_medio" decimal,
		"ticket_ultima_compra" decimal,
		"loja_mais_frequente" varchar(20),
		"loja_ultima_compra" varchar(20),
		"invalido" boolean
	  );`

	result, err := db.Exec(create)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.RowsAffected())
}
