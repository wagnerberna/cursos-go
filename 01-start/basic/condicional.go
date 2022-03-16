package basic

import "fmt"

func Condicional() {
	fmt.Println("---condicional:::")
	// Operadores:
	// Relacionais: == != >=
	// Lógicos: && ||
	// Negação: !
	fmt.Println("Operadores Relacionais")
	fmt.Println(1 == 1 && 1 > 1)
	fmt.Println(1 == 1 || 1 > 1)
	fmt.Println(!(1 != 1))
	// incrementais ++ --
	count := 0
	count++
	count += 10
	fmt.Println("Count:::", count)

	// Têm de ser entre {} e em linhas separadas
	if count < 10 {
		fmt.Println("menor q dez")
	} else {
		fmt.Println("maior q dez")
	}
}
