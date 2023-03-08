package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
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
