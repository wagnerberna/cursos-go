package basic

import "fmt"

func CondicionalIfInit() {
	fmt.Println("---condicional if init:::")
	// if init inicializa "declara" uma varíavel q não existia
	// já atribui um valor e faz o teste da condição usando a criada
	// Não acessa de fora do escopo
	num1 := 10
	if num2 := num1; num2 > 0 {
		println("maior q zero")
	}
}
