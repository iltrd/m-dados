package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

type Record struct {
	CPF             string `validate:"required,cpfcnpj"`
	Private         string `validate:"required"`
	Incompleto      string `validate:"required"`
	UltimaCompra    string `validate:"required"`
	TicketMedio     string `validate:"required"`
	TicketUltCompra string `validate:"required"`
	LojaFrequente   string `validate:"required"`
	LojaUltCompra   string `validate:"required"`
}

func main() {
	InsertData(db, data)
	CleanData(db, data)

	// Abrir arquivo
	f, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Ler arquivo CSV
	reader := csv.NewReader(f)
	reader.Comma = '|'
	reader.LazyQuotes = true

	var data [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, record)
	}

	// Conectar ao banco de dados
	db, err := sql.Open("postgres", "postgres://user:password@localhost/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inserir dados no banco de dados
	err = InsertData(db, CleanData(data))
	if err != nil {
		log.Fatal(err)
	}

	// Validar CPFs/CNPJs
	err = ValidateRecords(data)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData(db *sql.DB, data [][]string) error {
	for _, record := range data {
		cpf := record[0]
		private := record[1]
		incompleto := record[2]
		ultimacompra := record[3]
		ticketmedio := record[4]
		ticketultcompra := record[5]
		lojafrequente := record[6]
		lojaultcompra := record[7]

		query := fmt.Sprintf("INSERT INTO mytable (cpf, private, incompleto, ultimacompra, ticketmedio, ticketultcompra, lojafrequente, lojaultcompra) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')",
			cpf, private, incompleto, ultimacompra, ticketmedio, ticketultcompra, lojafrequente, lojaultcompra)
		_, err := db.Exec(query)

		if err != nil {
			return err
		}
	}

	return nil
}

func CleanData(data [][]string) [][]string {
	var cleanedData [][]string

	for _, record := range data {
		cleanedRecord := make([]string, len(record))
		for i, value := range record {
			cleanedValue := strings.Map(func(r rune) rune {
				if unicode.Is(unicode.Mn, r) {
					return -1
				}
				return unicode.ToLower(r)
			}, value)

			cleanedRecord[i] = cleanedValue
		}

		cleanedData = append(cleanedData, cleanedRecord)
	}

	return cleanedData
}

func ValidateRecords(data [][]string) error {
	validate := validator.New()

	for _, record := range data {
		r := Record{
			CPF:             record[0],
			Private:         record[1],
			Incompleto:      record[2],
			UltimaCompra:    record[3],
			TicketMedio:     record[4],
			TicketUltCompra: record[5],
			LojaFrequente:   record[6],
			LojaUltCompra:   record[7],
		}
		err := validate.Struct(r)
		if err != nil {
			return err
		}
	}

	return nil
}
