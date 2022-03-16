package class

import "fmt"

type pessoa struct {
	name string
	age  int8
}

type estudante struct {
	pessoa
	curso string
}

func Heranca(name string, age int8, curso string) {
	fmt.Println("Herança")
	p1 := pessoa{name, age}
	e1 := estudante{p1, curso}
	fmt.Println(e1)
	fmt.Println(e1.age)
}
