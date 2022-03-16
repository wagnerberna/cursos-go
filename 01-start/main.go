package main

import (
	"fmt"
	// pasta com mesmo nome do módulo (package)
	"modulo/basic"
	"modulo/class"
	"modulo/service"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Arquivo: Main")
	checkMail1 := checkmail.ValidateFormat("wag@gmail.com") //nil nulo ok
	checkMail2 := checkmail.ValidateFormat("wag.com")
	fmt.Println(checkMail1, checkMail2)

	basic.MessageDefault()
	basic.MessageCustom()
	basic.TypesVar()
	basic.Condicional()
	basic.CondicionalIfInit()
	basic.Loops()
	basic.Arrays()
	basic.Slices()
	basic.Maps()
	basic.ErroInterno()

	// atribuindo FN em uma var
	// declara ambos param. como string
	// podem ser declarados de forma individual
	var textPrint2 = func(text1, text2 string) string {
		return text1 + "::: " + text2
	}

	result1 := textPrint2("Texto FN textPrint2", "blablabla")
	fmt.Println(result1)
	// Fim

	println("---Service:::")
	println(service.Sum(2, 3))
	println(service.TextPrint("txt Arquivo service/text", "kkkkkk"))

	println(service.DoisRetornos(5, 2, "FN Dois retornos::: Sua soma é de: "))

	// duas var p/ receber o resultado de dois retornos
	resultText, resultSum := service.DoisRetornos(5, 2, "FN Dois retornos::: Sua soma é de: ")
	println(resultText, resultSum)

	// "_" Ignorar 1 dos resultados, mas recebe os dois retornos
	_, resultSum2 := service.DoisRetornos(5, 2, "FN Dois retornos::: Sua soma é de: ")
	println(resultSum2)

	// classe - struct
	class.Classe("Wagner", 40)
	// Herança
	class.Heranca("Wagner", 40, "TI")

}
