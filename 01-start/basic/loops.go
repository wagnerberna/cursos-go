package basic

import "fmt"

func Loops() {
	fmt.Println("--- Loops For:::")
	// só existe o for
	// forma simples
	inc := 0
	for inc <= 10 {
		fmt.Println("for1:", inc)
		inc += 2
	}
	// Forma Completa
	for count := 1; count <= 3; count++ {
		fmt.Println("for2:", count)
		inc++
	}

	// for Array / Slice
	nomes := []string{"wag", "Ali", "Cris"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// For Map
	// For Struct: Não funciona
	usuario := map[string]string{
		"nome":    "wagner",
		"apelido": "Fit",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, ":", valor)
	}

	// Loop infinito
	// for {
	// 	fmt.Println("loop eterno")
	// }
}
