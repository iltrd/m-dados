package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

type Record struct {
	CPF             string  `validate:"required,cpfcnpj"`
	Private         bool    `validate:"required"`
	Incompleto      bool    `validate:"required"`
	UltimaCompra    string  `validate:"required"`
	TicketMedio     float64 `validate:"required, decimal13"`
	TicketUltCompra float64 `validate:"required, decimal13"`
	LojaFrequente   string  `validate:"required"`
	LojaUltCompra   string  `validate:"required"`
}

func Decimal3(fl validator.FieldLevel) bool {
	// Regular expression to match decimal values with up to 3 decimal places
	decimalRegex := regexp.MustCompile(`^-?\d+(?:\.\d{1,3})?$`)

	// Check if the value matches the decimal regex
	value := fl.Field().String()
	if !decimalRegex.MatchString(value) {
		return false
	}

	return true
}

func CleanData(data [][]string) [][]string {
	var cleanedData [][]string

	for _, record := range data {
		cleanedRecord := make([]string, len(record))
		for i, field := range record {
			cleanedField := strings.TrimSpace(field)
			cleanedRecord[i] = cleanedField
		}
		cleanedData = append(cleanedData, cleanedRecord)
	}

	return cleanedData
}

func parseBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("Invalid boolean value: %s", s)
	}
}

func parseBoolOrPanic(s string) bool {
	b, err := parseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}

func ValidateRecords(data [][]string) error {
	validate := validator.New()

	for _, record := range data {
		ticketMedio, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		ticketUltCompra, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatal(err)
		}
		r := Record{
			CPF:             record[0],
			Private:         parseBoolOrPanic(record[1]),
			Incompleto:      parseBoolOrPanic(record[2]),
			UltimaCompra:    record[3],
			TicketMedio:     ticketMedio,
			TicketUltCompra: ticketUltCompra,
			LojaFrequente:   record[6],
			LojaUltCompra:   record[7],
		}

		err = validate.Struct(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func CalculateTicketMedio(data [][]string) (float64, error) {
	var total float64
	var count int

	for _, record := range data {
		ticketMedio, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return 0, err
		}

		total += ticketMedio
		count++
	}

	if count == 0 {
		return 0, nil
	}

	return total / float64(count), nil
}

func main() {
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

	// Calcular ticket médio
	ticketMedio, err := CalculateTicketMedio(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ticket médio: %.2f\n", ticketMedio)
}

func InsertData(db *sql.DB, data [][]string) error {
	// Preparar a instrução SQL para a inserção
	stmt, err := db.Prepare(`INSERT INTO record(cpf, private, incompleto, ultima_compra, ticket_medio, ticket_ult_compra, loja_frequente, loja_ult_compra) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Iterar sobre cada registro e inseri-lo no banco de dados
	for _, record := range data {
		if len(record) < 5 {
			fmt.Println("Record does not have enough fields:", record)
			continue
		}

		// Converter os valores para o tipo correto
		cpf := record[0]
		private := parseBoolOrPanic(record[1])
		incompleto := parseBoolOrPanic(record[2])
		ultimaCompra := record[3]
		ticketMedio, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatalf("Não foi possível converter o valor %s para float64. Erro: %v", record[4], err)
		}
		ticketUltCompra, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			log.Fatalf("Não foi possível converter o valor %s para float64. Erro: %v", record[5], err)
		}
		lojaFrequente := record[6]
		lojaUltCompra := record[7]

		// Executar a instrução SQL de inserção
		_, err = stmt.Exec(cpf, private, incompleto, ultimaCompra, ticketMedio, ticketUltCompra, lojaFrequente, lojaUltCompra)
		if err != nil {
			return err
		}
	}

	return nil
}
