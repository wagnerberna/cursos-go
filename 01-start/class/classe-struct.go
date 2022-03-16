package class

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  uint8
}

func Classe(name string, age uint8) {
	fmt.Println("Classes: Structs")

	// 1ª forma cria a instância depois passa os valores
	var newUser1 user
	newUser1.name = name
	newUser1.age = age
	fmt.Println(newUser1)

	//  2ª forma cria e a instância e já passa os valores
	newUser2 := user{name, age}
	fmt.Println(newUser2)

	//  3ª forma argumento nomeado sem algum campo
	newUser3 := user{name: name}
	fmt.Println(newUser3)

	fmt.Println("TIPO STRUCT:", reflect.TypeOf(newUser1))

	
}
