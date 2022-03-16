package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	// "encoding/csv"
	"strings"
	// "regexp"
	"reflect"
	"moduleFile/service"
)
// fmt - implementando a funcionalidade de E/S formatada
// os - para uma interface independente de plataforma para a funcionalidade do sistema operacional
// log - pacote padrão de extração
// bufio - suporta E/S tamponadas

func main() {
	fmt.Println("Manipulação Arquilo txt/csv")

	const fileName string = "base_teste.txt" 
	// abrir o arquivo  retorna um tipo de ponteiro para o arquivo.
	file, err := os.Open(fileName)
	// Tratamento de erros ao abrir o arquivo
	if err !=nil {
		log.Fatalf("Error opening file: %s", err)
	}

	//  retorna um tipo de scanner que possui funções que suportam a leitura através do arquivo.
	fileScanner := bufio.NewScanner(file)

	// Dois métodos para ler linha por linha no arquivo:
	// Scanner que avança o scanner para o nova linha
	// Text que lê a linha mais recente que foi gerado quando o Scan foi chamado, retorna uma string

	for fileScanner.Scan(){
		lineString := fileScanner.Text()
		// fmt.Println("Ponto0:")
		// fmt.Println(lineString)
		// fmt.Println(reflect.TypeOf(lineString))

		// strings pacote tem um Fields método p/ remover espaços em branco extras
		// converte para um slice de strings
		lineFields := strings.Fields(lineString)
		// fmt.Println("Ponto1:")
		// fmt.Println(lineFields)

		fmt.Println(reflect.TypeOf(lineFields))
		// fmt.Println(len(lineFields))
		// fmt.Println("Campos:::", lineFields[0], lineFields[3], lineFields[4], lineFields[7] )

		payload := service.DataPayload(lineFields)
		fmt.Println("---Data Type and Payload:::")
		fmt.Println(reflect.TypeOf(payload))
		fmt.Println(payload)

		// TODOS CAMPOS:
			// cpf := lineFields[0]
			// private := lineFields[1]
			// incompleto := lineFields[2]
			// dateLastPurchase := lineFields[3]
			// ticketAverage := lineFields[4]
			// ticketLastPurchase := lineFields[5]
			// cnpjStoreMostFrequent := lineFields[6]
			// cnpjStoreLastPurchase := lineFields[7]
	
		// Tratamento de erro durante a leitura
		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error reading file: %s", err)
		}
	}
	// Fechar o arquivo TESTAR
	// file.Closed()
}
