package service

import (
	"fmt"
	"strconv"
	"strings"
	// "math"
)

type Payload struct {
	cpf string
	private bool
	incomplete bool
	dateLastPurchase string
	ticketAverage float64
	ticketLastPurchase float64
	cnpjStoreMostFrequent string
	cnpjStoreLastPurchase string
}

func CleanCpfCnpj(value string) string {
	// fmt.Println("Clean CNPJ / CPF")
	if value == "NULL" {
		return ""
	}
	cleanValue := strings.ReplaceAll(value, ".", "")
	cleanValue = strings.ReplaceAll(cleanValue, "/", "")
	cleanValue = strings.ReplaceAll(cleanValue, "-", "")
	return cleanValue
}

// func CheckCpf() {
// 	fmt.Println("CheckCpf")
// }

// func CheckCnpj() {
// 	fmt.Println("CheckCnpj")
// }

func CleanTicketValue(valueStr string) float64 {
	// fmt.Println("CleanValue:", valueStr)
	if valueStr == "NULL" {
		return 0
	}
	cleanValueStr := strings.Replace(valueStr, ",", ".", 1)
	// println("cleanValueStr:", cleanValueStr)

	convertValueStrToFloat, _ := strconv.ParseFloat(cleanValueStr, 64)

	return convertValueStrToFloat
	
}

func ConvertBoolean(value string) bool {
	if value == "1" {
		return true
	}
	return false
}

func ConvertDate(date string) string {
	if date == "NULL" {
		return ""
	}
	cleanDate := strings.ReplaceAll(date, "-", "")
	// fmt.Printf(cleanDate)
	
	return cleanDate
} 

func DataPayload(data []string) Payload {
	// fmt.Println("-DataClean:::")
	// fmt.Println(data)
	// fmt.Println("Campos 0 4 7", data[0], data[4], data[7])

	// replace
	// strings.Replace(Originalstr, oldstr, newstr, num. max de trocas int)

	// TODOS CAMPOS:
	cpf := CleanCpfCnpj(data[0])
	private := ConvertBoolean(data[1])
	incomplete := ConvertBoolean(data[2])
	dateLastPurchase := ConvertDate(data[3])
	ticketAverage := CleanTicketValue(data[4])
	ticketLastPurchase := CleanTicketValue(data[5])
	cnpjStoreMostFrequent := CleanCpfCnpj(data[6])
	cnpjStoreLastPurchase := CleanCpfCnpj(data[7])

	// fmt.Println("ticketAverage and ticketLastPurchase Result:::", ticketAverage, ticketLastPurchase)
	// fmt.Println("cpf, cnpjStoreMostFrequent, cnpjStoreLastPurchase:::", cpf, cnpjStoreMostFrequent, cnpjStoreLastPurchase)
	// fmt.Println("private, incompete:::", private, incomplete)
	// fmt.Println("dateLastPurchase:::", dateLastPurchase)

	// TODOS CAMPOS STRUCT:
	payload := Payload{cpf, private, incomplete, dateLastPurchase, ticketAverage, ticketLastPurchase, cnpjStoreMostFrequent, cnpjStoreLastPurchase}

	// fmt.Println("---Data Payload:::")
	// fmt.Println(dataPayload)

	return payload

	
}