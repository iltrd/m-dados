package main

import (
	"fmt"

	_ "github.com/iltrd/manipular-dados/models"
	_ "github.com/iltrd/manipular-dados/utils"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1404"
	dbname   = "postgres"
)

func main() {
	// conecta ao banco de dados
	db, err := utils.Connect(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// cria tabela customers
	err = utils.CreateTable(db)
	if err != nil {
		panic(err)
	}

	// lê arquivo csv/txt
	customers, err := utils.ReadCsv("clientes.csv")
	if err != nil {
		panic(err)
	}

	// insere dados no banco de dados
	err = utils.InsertData(db, customers)
	if err != nil {
		panic(err)
	}

	fmt.Println("Processo concluído com sucesso")
}
