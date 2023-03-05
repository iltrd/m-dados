package utils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/iltrd/manipular-dados/models"
)

// lê arquivo csv/txt e retorna slice de Customer
func ReadCsv(filename string) ([]models.Customer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t' // define separador como tabulação

	var customers []models.Customer
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		cpf, _ := strconv.Atoi(line[0])
		private, _ := strconv.Atoi(line[1])
		incompleto, _ := strconv.Atoi(line[2])
		ultimaCompra, _ := time.Parse("2006-01-02", line[3])
		ticketMedio, _ := strconv.ParseFloat(line[4], 64)
		ticketUltComp, _ := strconv.ParseFloat(line[5], 64)
		lojaFrequente, _ := strconv.Atoi(line[6])
		lojaUltComp, _ := strconv.Atoi(line[7])

		customer := models.Customer{
			Cpf:           cpf,
			Private:       private,
			Incompleto:    incompleto,
			UltimaCompra:  ultimaCompra,
			TicketMedio:   ticketMedio,
			TicketUltComp: ticketUltComp,
			LojaFrequente: lojaFrequente,
			LojaUltComp:   lojaUltComp,
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
