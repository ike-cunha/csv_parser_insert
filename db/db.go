package db

import "time"

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
