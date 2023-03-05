package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// conexão com o banco de dados
func Connect() (*sql.DB, error) {
	connStr := "postgres://postgres:password@db:5432/customers?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database")
	return db, nil
}

// cria tabela no banco de dados
func CreateTable(db *sql.DB) error {
	sql := `
		CREATE TABLE
		customers (
			cpf INTEGER PRIMARY KEY,
			private INTEGER,
			incompleto INTEGER,
			ultima_compra TIMESTAMP,
			ticket_medio FLOAT,
			ticket_ult_comp FLOAT,
			loja_frequente INTEGER,
			loja_ult_comp INTEGER
			)
			`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println("Table customers created")
	return nil
}

// insere dados no banco de dados
func InsertData(db *sql.DB, customers []Customer) error {
	sql := "INSERT INTO customers ( cpf, private, incompleto, ultima_compra, ticket_medio, ticket_ult_comp, loja_frequente, loja_ult_comp ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	for _, c := range customers {
		_, err := db.Exec(sql, c.Cpf, c.Private, c.Incompleto, c.UltimaCompra, c.TicketMedio, c.TicketUltComp, c.LojaFrequente, c.LojaUltComp)
		if err != nil {
			return err
		}
	}

	fmt.Println("Data inserted successfully")
	return nil
}
