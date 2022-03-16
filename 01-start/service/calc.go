package service

// após() tipagem do tipo do retorno
func Sum(num1, num2 int8) int8 {
	sum := num1 + num2
	return sum
}

// Go permite uma FN ter dois retornos
// declarar tipo de retorno na mesma ordem dos valores do return
func DoisRetornos(num1 int8, num2 int8, text string) (string, int8) {
	message := text
	sum := Sum(num1, num2)
	return message, sum
}
