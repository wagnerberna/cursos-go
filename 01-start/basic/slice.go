package basic

import (
	"fmt"
	"reflect"
)

func Slices() {
	fmt.Println("---Slices:::")
	// Não é um array o tipo é Slice "Fatia"
	// se declara da mesma forma q um array mas sem passar o tamanho
	// Não têm tamanho fixo
	slice1 := []int{21, 22, 23}

	// Reflect mostra o tipo
	fmt.Println("TIPO SLICE", reflect.TypeOf(slice1))
	// len / cap (Tamanho e capacidade)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// append add item e retorna um slice novo
	slice1 = append(slice1, 24)
	fmt.Println(slice1)

	// usa índices intervalos igual python
	slice2 := slice1[1:2]
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("---Make / Slices::: ???")
}
