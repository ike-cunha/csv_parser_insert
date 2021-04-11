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
