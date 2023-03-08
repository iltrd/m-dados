package main

import (
	"github.com/go-playground/validator"
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

func ValidateRecords(records [][]string) error {
	validate := validator.New()
	for _, record := range records {
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
