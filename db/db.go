package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/klassmann/cpfcnpj"
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

//Inserts the received file into the database
func Insert(file []byte) {
	db, err := sql.Open("postgres", CONNECTION_STRING)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// CREATE table
	createTable(db)

	var store Store
	data := strings.Split(string(file), "\n")[1:]
	for _, values := range data {
		value := strings.Fields(values)

		store.CPF = value[0]
		store.Privado = stringToBool(value[1])
		store.Incompleto = stringToBool(value[2])
		store.DataUltimaCompra = stringToDate(value[3])
		store.TicketMedio = stringToFloat(value[4])
		store.TicketUltimaCompra = stringToFloat(value[5])
		store.LojaMaisFrequente = value[6]
		store.LojaUltimaCompra = value[7]
		store.Invalido = cpfCnpjValidation(store.CPF, store.LojaMaisFrequente, store.LojaUltimaCompra)

		insert := `
				INSERT INTO store(
					cpf,
					privado,
					incompleto,
					data_ultima_compra,
					ticket_medio,
					ticket_ultima_compra,
					loja_mais_frequente,
					loja_ultima_compra,
					invalido
				) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`

		_, err = db.Exec(insert,
			store.CPF,
			store.Privado,
			store.Incompleto,
			store.DataUltimaCompra,
			store.TicketMedio,
			store.TicketUltimaCompra,
			store.LojaMaisFrequente,
			store.LojaUltimaCompra,
			store.Invalido,
		)

		if err != nil {
			panic(err)
		}
	}
}

//Parses string to boolean
func stringToBool(text string) bool {
	result, err := strconv.ParseBool(text)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//Parses string to float
func stringToFloat(text string) float64 {
	if text == "NULL" {
		return 0
	}

	price, err := strconv.ParseFloat(strings.Replace(text, ",", ".", 1), 64)
	if err != nil {
		log.Fatal(err)
	}
	return price
}

//Parses string to date
func stringToDate(date_text string) time.Time {
	if date_text == "NULL" {
		return time.Time{}
	}

	date, err := time.Parse("2006-01-02", date_text)
	if err != nil {
		log.Fatal(err)
	}
	return date
}

//Validates registration numbers: CPF and CNPJ
func cpfCnpjValidation(cpf string, cnpjFrequente string, cnpjUltimaCompra string) bool {
	if cpf == "NULL" || cnpjFrequente == "NULL" || cnpjUltimaCompra == "NULL" {
		return true
	}

	cpfIsValid := cpfcnpj.ValidateCPF(cpf)
	cnpjFrequenteIsValid := cpfcnpj.ValidateCNPJ(cnpjFrequente)
	cnpjUltimaCompraIsValid := cpfcnpj.ValidateCNPJ(cnpjUltimaCompra)

	return cpfIsValid && cnpjFrequenteIsValid && cnpjUltimaCompraIsValid
}
