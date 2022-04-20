package controller

// !!!!PACOTE TODO MIGRADO PARA DENTRO DO MAIN!!!
// PARA SER MAIS RÁPIDO CONEXÃO CRIADO FORA DO FOR

import (
	"database/sql"
	"fmt"
	"moduleFile/service"
	"strconv"

	// "reflect"

	_ "github.com/lib/pq"
)

// FN DotEnv
// func goDotEnvVariable(key string) string {
// 	// load .env file
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		panic(err)
// 	}

// 	return os.Getenv(key)
// }

// método direto de declaração:
// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123"
// 	dbname   = "neoway"
// )

// fmt.Sprintf combina strings
func AddFieldsDb(lineFields []string) {
	// Declaração .env
	host := service.DotEnvVariable("DB_HOST")
	port, _ := strconv.Atoi(service.DotEnvVariable("DB_PORT"))
	user := service.DotEnvVariable("DB_USER")
	password := service.DotEnvVariable("DB_PASSWORD")
	dbname := service.DotEnvVariable("D_BNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Testa a conexão usando um ping:
	// err = db.Ping()
	// if err != nil {
	// panic(err)
	// }
	// fmt.Println("Successfully connected!")
	// Fim Teste conexão

	// fmt.Println("Controller Data:::", lineFields)

	payload := service.DataPayload(lineFields)

	// fmt.Println("Controller payload:::", payload)
	// fmt.Println("payload:::", payload)
	// fmt.Println("Type payload controller:::")
	// fmt.Println(reflect.TypeOf(payload))
	// print struct com o nome das chaves
	// fmt.Println("payload print chave e valor:::")
	// fmt.Printf("%+v\n",payload)
	// fmt.Println("payload print completo:::")
	// fmt.Printf("%#v",payload)

	// acessar campo payload:
	// campo com 1ª Letra maiúscula
	// fmt.Println("campo payload:::", payload.Cpf)
	// fmt.Println(reflect.TypeOf(payload.Cpf))
	// fmt.Println("campo payload:::", payload.Incomplete)
	// fmt.Println(reflect.TypeOf(payload.Incomplete))

	// "RETURNING customer_id" dentro das crases para retornar o id
	sqlStatement := `
	INSERT INTO public.customer (cpf, private, incomplete, date_last_purchase, ticket_average, ticket_last_purchase, cnpj_store_most_frequent, cnpj_store_last_purchase)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	// RETURNING customer_id`

	// declarar a variável para retornar o id
	// customer_id := 0
	// ".Scan(&customer_id)" no final da query para retornar o id
	// err = db.QueryRow(sqlStatement, payload.Cpf, payload.Private, payload.Incomplete, payload.DateLastPurchase, payload.TicketAverage, payload.TicketLastPurchase, payload.CnpjStoreMostFrequent, payload.CnpjStoreLastPurchase).Scan(&customer_id)

	db.QueryRow(sqlStatement, payload.Cpf, payload.Private, payload.Incomplete, payload.DateLastPurchase, payload.TicketAverage, payload.TicketLastPurchase, payload.CnpjStoreMostFrequent, payload.CnpjStoreLastPurchase).Scan()
	// if err != nil {
	// 	fmt.Println("Create record:", err)
	// }

	// Retorno do id
	// fmt.Println("New record ID is:", customer_id)

}
