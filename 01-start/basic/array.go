package basic

import (
	"fmt"
	"reflect"
)


func Arrays() {
	fmt.Println("---Arrays:::")
	// declaração passando num. de el e tipo ambos fixos
	// não existe array de posições dinâmicas se usa mais o Slice
	// cria um array de 5 zeros ou espaços vazios se for string
	// só aceita el do mesmo tipo
	var array1 [5]int
	var array2 [5]string
	fmt.Println(array1, array2)
	array1[2] = 10
	array2[4] = "hello"

	fmt.Println(array1, array2)
	fmt.Println("TIPO ARRAY:", reflect.TypeOf(array1))

	// atribuindo valor na criação
	array3 := [5]int{1, 2, 3, 4, 5}
	array4 := [3]string{"a", "b", "c"}
	// "..." array onde o num. de posições vai ser definido pelo num. de el
	array5 := [...]int{11, 12, 13}
	fmt.Println(array3, array4, array5)

}
