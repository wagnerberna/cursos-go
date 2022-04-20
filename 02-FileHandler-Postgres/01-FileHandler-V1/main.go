package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "encoding/csv"
	"strings"
	// "regexp"

	// "moduleFile/service"

	"moduleFile/service"
)

// fmt - implementando a funcionalidade de E/S formatada
// os - para uma interface independente de plataforma para a funcionalidade do sistema operacional
// log - pacote padrão de extração
// bufio - suporta E/S tamponadas

func main() {
	service.CreateDatabase()
	dbname := service.DotEnvVariable("D_BNAME")
	// fmt.Println("Manipulação Arquilo txt/csv")
	fmt.Println("Process started.")

	const fileName string = "base_teste.txt"
	// abrir o arquivo  retorna um tipo de ponteiro para o arquivo.
	file, err := os.Open(fileName)
	// Tratamento de erros ao abrir o arquivo
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	//  retorna um tipo de scanner que possui funções que suportam a leitura através do arquivo.
	fileScanner := bufio.NewScanner(file)

	db := service.ConnectDb(dbname)
	// Dois métodos para ler linha por linha no arquivo:
	// Scanner que avança o scanner para o nova linha
	// Text que lê a linha mais recente que foi gerado quando o Scan foi chamado, retorna uma string

	for fileScanner.Scan() {
		lineString := fileScanner.Text()
		// fmt.Println("Ponto0:")
		// fmt.Println(lineString)
		// fmt.Println(reflect.TypeOf(lineString))

		// strings pacote tem um Fields método p/ remover espaços em branco extras
		// converte para um slice de strings
		lineFields := strings.Fields(lineString)
		// fmt.Println("Ponto1:")
		// fmt.Println(lineFields)

		// fmt.Println(reflect.TypeOf(lineFields))
		// fmt.Println(len(lineFields))
		// fmt.Println("Campos:::", lineFields[0], lineFields[3], lineFields[4], lineFields[7] )

		// TESTES passar para o controller:
		// payload := service.DataPayload(lineFields)
		// fmt.Println("---Data Type and Payload:::")
		// fmt.Println(reflect.TypeOf(payload))
		// fmt.Println(payload)

		checkFirstLine := lineFields[0] == "CPF"
		// fmt.Println(checkFirstLine)
		if !checkFirstLine {
			// controller.AddFieldsDb(lineFields)
			payload := service.DataPayload(lineFields)
			sqlAddCustomer := `
			INSERT INTO public.customer (cpf, private, incomplete, date_last_purchase, ticket_average, ticket_last_purchase, cnpj_store_most_frequent, cnpj_store_last_purchase)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			db.QueryRow(sqlAddCustomer, payload.Cpf, payload.Private, payload.Incomplete, payload.DateLastPurchase, payload.TicketAverage, payload.TicketLastPurchase, payload.CnpjStoreMostFrequent, payload.CnpjStoreLastPurchase).Scan()

		}

		// Tratamento de erro durante a leitura
		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error reading file: %s", err)
		}
	}

	count := 0

	db.QueryRow("SELECT COUNT(*) FROM customer").Scan(&count)
	// Fechar o arquivo TESTAR
	// file.Closed()

	defer db.Close()

	fmt.Println("Process finished, total records: ", count)

}
