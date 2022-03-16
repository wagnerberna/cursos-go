package basic

import (
	"errors"
	"fmt"
)

func ErroInterno() {
	// Erro necessário criar tipo, usa pacote nativo "errors"
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
