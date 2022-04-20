package service

import (
	// "fmt"
	"strconv"
	"strings"
)

type Payload struct {
	Cpf                   string
	Private               bool
	Incomplete            bool
	DateLastPurchase      string
	TicketAverage         float64
	TicketLastPurchase    float64
	CnpjStoreMostFrequent string
	CnpjStoreLastPurchase string
}

func CleanCpfCnpj(value string) string {
	if value == "NULL" {
		return ""
	}
	cleanValue := strings.ReplaceAll(value, ".", "")
	cleanValue = strings.ReplaceAll(cleanValue, "/", "")
	cleanValue = strings.ReplaceAll(cleanValue, "-", "")
	return cleanValue
}

func CleanTicketValue(valueStr string) float64 {
	if valueStr == "NULL" {
		return 0
	}
	cleanValueStr := strings.Replace(valueStr, ",", ".", 1)
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

	return cleanDate
}

func DataPayload(data []string) Payload {
	cpf := CleanCpfCnpj(data[0])
	private := ConvertBoolean(data[1])
	incomplete := ConvertBoolean(data[2])
	dateLastPurchase := ConvertDate(data[3])
	ticketAverage := CleanTicketValue(data[4])
	ticketLastPurchase := CleanTicketValue(data[5])
	cnpjStoreMostFrequent := CleanCpfCnpj(data[6])
	cnpjStoreLastPurchase := CleanCpfCnpj(data[7])

	return Payload{
		cpf, private, incomplete, dateLastPurchase, ticketAverage,
		ticketLastPurchase, cnpjStoreMostFrequent, cnpjStoreLastPurchase,
	}
}
