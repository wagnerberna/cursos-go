package basic

import (
	"fmt"
)

func TypesVar() {
	// Tipos de atribuições:
	// "=" Tem de declarar o tipo
	// := Não precisa declarar o tipo usa inferência

	// Tipos de dados
	// int8 / 16 / 32 / 64 ou apenas int  q pega do processador 64 bits
	// dif de tamanho de valor suportado
	const num1 int = 100
	num2 := 100
	// int sem sinal não aceita negativo
	const num3 uint = 2

	// float8 / 16 / 32 / 64 (não aceita apenas float)
	const num4 float64 = 1050.50
	const num5 float64 = 2050.35

	fmt.Println("Soma:::", num1+num2+int(num3))
	fmt.Println("subtração:::", num4-num5)

	// String
	text1 := "Bom dia:"

	// 1 caracter vai aparecer o valor na tabela ascii se usar aspas simples
	textChar1 := "b"
	// só aceita 1 caracter com aspas simples
	textChar2 := 'b'
	const textChar3 string = "b"

	fmt.Println("String:::", text1, textChar1, textChar2, textChar3)

	//bolean
	const test1 bool = false
	const test2 bool = true

	fmt.Println("Booleano:::", test1, test2)

}
