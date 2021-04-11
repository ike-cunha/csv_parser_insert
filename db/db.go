// Package db provides a database instance, a table struct
//and some CRUD func.
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

type Purchase struct {
	CPF                      string
	Privado                  bool
	Incompleto               bool
	DataUltimaCompra         time.Time
	TicketMedio              float64
	TicketUltimaCompra       float64
	LojaMaisFrequente        string
	LojaUltimaCompra         string
	DadosCadastraisInvalidos bool
}

//Creates Purchase table
func createTable(db *sql.DB) {
	create := `CREATE TABLE IF NOT EXISTS "purchase"  (
		"id" SERIAL PRIMARY KEY,
		"cpf" varchar(20),
		"privado" boolean,
		"incompleto" boolean,
		"data_ultima_compra" date,
		"ticket_medio" decimal,
		"ticket_ultima_compra" decimal,
		"loja_mais_frequente" varchar(20),
		"loja_ultima_compra" varchar(20),
		"dados_cadastrais_invalidos" boolean
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

	var purchase Purchase
	data := strings.Split(string(file), "\n")[1:]
	for _, values := range data {
		value := strings.Fields(values)

		purchase.CPF = value[0]
		purchase.Privado = stringToBool(value[1])
		purchase.Incompleto = stringToBool(value[2])
		purchase.DataUltimaCompra = stringToDate(value[3])
		purchase.TicketMedio = stringToFloat(value[4])
		purchase.TicketUltimaCompra = stringToFloat(value[5])
		purchase.LojaMaisFrequente = value[6]
		purchase.LojaUltimaCompra = value[7]
		purchase.DadosCadastraisInvalidos = cpfCnpjValidation(purchase.CPF, purchase.LojaMaisFrequente, purchase.LojaUltimaCompra)

		insert := `
				INSERT INTO purchase(
					cpf,
					privado,
					incompleto,
					data_ultima_compra,
					ticket_medio,
					ticket_ultima_compra,
					loja_mais_frequente,
					loja_ultima_compra,
					dados_cadastrais_invalidos
				) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
					RETURNING id`

		var id int
		row := db.QueryRow(insert,
			purchase.CPF,
			purchase.Privado,
			purchase.Incompleto,
			purchase.DataUltimaCompra,
			purchase.TicketMedio,
			purchase.TicketUltimaCompra,
			purchase.LojaMaisFrequente,
			purchase.LojaUltimaCompra,
			purchase.DadosCadastraisInvalidos,
		).Scan(&id)

		if row != nil {
			panic(err)
		}

		sanitize(value, int64(id), db)
	}
}

//Removes special characters from the cpf, loja_mais_frequente, and loja_ultima_compra columns
func sanitize(value []string, id int64, db *sql.DB) {
	var purchase Purchase

	purchase.CPF = value[0]
	purchase.LojaMaisFrequente = value[6]
	purchase.LojaUltimaCompra = value[7]

	update := `
			UPDATE purchase
			SET cpf=regexp_replace($1, '[^a-zA-Z0-9]+', '','g'),
				loja_mais_frequente=regexp_replace($2, '[^a-zA-Z0-9]+', '','g'),
				loja_ultima_compra=regexp_replace($3, '[^a-zA-Z0-9]+', '','g')
			WHERE id=$4;`

	_, err := db.Exec(update,
		purchase.CPF,
		purchase.LojaMaisFrequente,
		purchase.LojaUltimaCompra,
		id,
	)

	if err != nil {
		panic(err)
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
