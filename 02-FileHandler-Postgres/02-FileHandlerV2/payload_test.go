package main

import (
	"moduleFile/service"
	"testing"
)

func TestPayloadFields(t *testing.T) {
	dataInput := []string{
		"345.244.728-60", "0", "1", "2013-06-12", "91,28", "91,28", "79.379.491/0008-50", "79.379.491/0008-50"}
	responseExpected := service.Payload{
		Cpf:                   "34524472860",
		Private:               false,
		Incomplete:            true,
		DateLastPurchase:      "20130612",
		TicketAverage:         91.28,
		TicketLastPurchase:    91.28,
		CnpjStoreMostFrequent: "79379491000850",
		CnpjStoreLastPurchase: "79379491000850",
	}
	responseResult := service.DataPayload(dataInput)

	if responseExpected != responseResult {
		t.Errorf(`Test service payload failed.
		Response expected: %v 
		Response result: %v`, responseExpected, responseResult)
	}
}

func TestPayloadFieldsNull(t *testing.T) {
	dataInput := []string{
		"345.244.728-60", "1", "0", "NULL", "NULL", "NULL", "NULL", "NULL"}
	responseExpected := service.Payload{
		Cpf:                   "34524472860",
		Private:               true,
		Incomplete:            false,
		DateLastPurchase:      "",
		TicketAverage:         0,
		TicketLastPurchase:    0,
		CnpjStoreMostFrequent: "",
		CnpjStoreLastPurchase: "",
	}
	responseResult := service.DataPayload(dataInput)

	if responseExpected != responseResult {
		t.Errorf(`Test service payload failed.
		Response expected: %v
		Response result: %v`, responseExpected, responseResult)
	}
}
