package main

import (
	"bufio"
	"fmt"
	"log"
	"moduleFile/service"
	"os"
	"strings"
)

func main() {
	service.CreateDatabase()
	fmt.Println("Digite o nome do arquivo:")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Println("Process started.")
	dbname := service.DotEnvVariable("D_BNAME")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	db := service.ConnectDb(dbname)

	for fileScanner.Scan() {
		lineString := fileScanner.Text()
		lineFields := strings.Fields(lineString)

		checkFirstLine := lineFields[0] == "CPF" || lineFields[0] == "cpf"

		fmt.Println(lineFields)
		if !checkFirstLine {
			payload := service.DataPayload(lineFields)
			fmt.Println(payload)

			sqlAddCustomer := `
			INSERT INTO public.customer (cpf, private, incomplete, date_last_purchase, ticket_average,
				 ticket_last_purchase, cnpj_store_most_frequent, cnpj_store_last_purchase)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			db.QueryRow(sqlAddCustomer, payload.Cpf, payload.Private, payload.Incomplete, payload.DateLastPurchase,
				payload.TicketAverage, payload.TicketLastPurchase, payload.CnpjStoreMostFrequent, payload.CnpjStoreLastPurchase).Scan()
		}

		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error reading file: %s", err)
		}
	}

	count := 0
	db.QueryRow("SELECT COUNT(*) FROM customer").Scan(&count)
	defer db.Close()

	fmt.Println("Process finished, total records: ", count)
}
