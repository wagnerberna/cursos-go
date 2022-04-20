package service

import (
	"database/sql"
	"fmt"

	"strconv"

	_ "github.com/lib/pq"
)

func ConnectDb(database string) *sql.DB {
	host := DotEnvVariable("DB_HOST")
	port, _ := strconv.Atoi(DotEnvVariable("DB_PORT"))
	user := DotEnvVariable("DB_USER")
	password := DotEnvVariable("DB_PASSWORD")
	dbname := DotEnvVariable("D_BNAME")

	psqlInfo := ""

	if database == dbname {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=postgres sslmode=disable",
			host, port, user, password)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateDatabase() {
	dbname := DotEnvVariable("D_BNAME")
	db := ConnectDb("")

	// psqlCreateDb := fmt.Sprintf("DROP DATABASE %s;", dbname)
	psqlCreateDb := fmt.Sprintf("CREATE DATABASE %s;", dbname)
	db.QueryRow(psqlCreateDb).Scan()
	defer db.Close()

	db = ConnectDb(dbname)
	sqlCreateTable := `
	CREATE TABLE public.customer (
	customer_id serial4 NOT NULL,
	cpf varchar NOT NULL,
	private bool NULL,
	incomplete bool NULL,
	date_last_purchase varchar NULL,
	ticket_average varchar NULL,
	ticket_last_purchase varchar NULL,
	cnpj_store_most_frequent varchar NULL,
	cnpj_store_last_purchase varchar NULL,
	CONSTRAINT customer_pk PRIMARY KEY (customer_id)
	);`

	db.QueryRow(sqlCreateTable).Scan()
	defer db.Close()
}
