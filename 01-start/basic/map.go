package basic

import (
	"fmt"
	"reflect"
)

func Maps() {
	fmt.Println("---Maps (dicionários):::")
	// semelhante ao dicionário ou objetos chave valor
	// map entre [] Todos campos do mesmo Tipo
	// struct entre {} permite valores tipos diferentes, mas usa instância
	// pode ser feito aninhado
	// [tipo chaves] tipo valores
	usuario := map[string]string{
		"nome":    "wagner",
		"apelido": "Fit",
	}
	fmt.Println(usuario["apelido"])
	delete(usuario, "nome")
	fmt.Println(usuario)
	fmt.Println("TIPO MAPS:", reflect.TypeOf(usuario))

}
