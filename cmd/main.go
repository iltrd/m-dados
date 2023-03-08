package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Cliente struct {
	Cpf             string
	Private         int
	Incompleto      int
	UltimaCompra    string
	TicketMedio     float64
	TicketUltCompra float64
	LojaFrequente   string
	LojaUltCompra   string
}

func main() {
	// Abrir o arquivo csv/txt
	file, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Configurar o banco de dados
	db, err := sql.Open("postgres", "postgres://testuser:testpassword@localhost/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar a tabela clientes se ela não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clientes (
		cpf VARCHAR(11) NOT NULL PRIMARY KEY,
		private INT NOT NULL,
		incompleto INT NOT NULL,
		ultimacompra DATE NOT NULL,
		ticketmedio FLOAT NOT NULL,
		ticketultcompra FLOAT NOT NULL,
		lojafrequente VARCHAR(100) NOT NULL,
		lojaultcompra VARCHAR(100) NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Preparar a query de inserção de clientes
	stmt, err := db.Prepare(`INSERT INTO clientes (
		cpf,
		private,
		incompleto,
		ultimacompra,
		ticketmedio,
		ticketultcompra,
		lojafrequente,
		lojaultcompra
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Ler e processar os dados do arquivo csv/txt
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	var clientes []Cliente
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		cpf := strings.ReplaceAll(record[0], ".", "")
		cpf = strings.ReplaceAll(cpf, "-", "")
		private, _ := strconv.Atoi(record[1])
		incompleto, _ := strconv.Atoi(record[2])
		ultimaCompra := record[3]
		ticketMedio, _ := strconv.ParseFloat(record[4], 64)
		ticketUltCompra, _ := strconv.ParseFloat(record[5], 64)
		lojaFrequente := record[6]
		lojaUltCompra := record[7]

		clientes = append(clientes, Cliente{
			Cpf:             cpf,
			Private:         private,
			Incompleto:      incompleto,
			UltimaCompra:    ultimaCompra,
			TicketMedio:     ticketMedio,
			TicketUltCompra: ticketUltCompra,
			LojaFrequente:   lojaFrequente,
			LojaUltCompra:   lojaUltCompra,
		})
	}

	// Inserir os clientes no banco de dados
	for _, c := range clientes {
		_, err := stmt.Exec(
			strings.ToUpper(c.Cpf),
			c.Private,
			c.Incompleto,
			c.UltimaCompra,
			c.TicketMedio,
			c.TicketUltCompra,
			strings.ToUpper(c.LojaFrequente),
			strings.ToUpper(c.LojaUltCompra),
		)
		if err != nil {
			log.Println(err)
		}
	}

	// Realizar a higienização dos dados e validar os CPFs
	rows, err := db.Query("SELECT * FROM clientes;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cliente Cliente
		err := rows.Scan(
			&cliente.Cpf,
			&cliente.Private,
			&cliente.Incompleto,
			&cliente.UltimaCompra,
			&cliente.TicketMedio,
			&cliente.TicketUltCompra,
			&cliente.LojaFrequente,
			&cliente.LojaUltCompra,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Higienizar CPF
		cpfClean := strings.ReplaceAll(cliente.Cpf, ".", "")
		cpfClean = strings.ReplaceAll(cpfClean, "-", "")
		cliente.Cpf = cpfClean

		// Validar CPF
		if !utils.IsCpfValid(cliente.Cpf) {
			log.Printf("CPF inválido: %s", cliente.Cpf)
		}

		// Higienizar lojas
		cliente.LojaFrequente = utils.RemoveSpecialChars(strings.ToUpper(cliente.LojaFrequente))
		cliente.LojaUltCompra = utils.RemoveSpecialChars(strings.ToUpper(cliente.LojaUltCompra))

		// Atualizar registro no banco de dados com os dados higienizados
		_, err = db.Exec(`UPDATE clientes SET 
			cpf=$1, 
			private=$2, 
			incompleto=$3, 
			ultimacompra=$4, 
			ticketmedio=$5, 
			ticketultcompra=$6, 
			lojafrequente=$7, 
			lojaultcompra=$8 
			WHERE cpf=$1;`,
			cliente.Cpf,
			cliente.Private,
			cliente.Incompleto,
			cliente.UltimaCompra,
			cliente.TicketMedio,
			cliente.TicketUltCompra,
			cliente.LojaFrequente,
			cliente.LojaUltCompra,
		)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Dados importados e tratados com sucesso!")

}
