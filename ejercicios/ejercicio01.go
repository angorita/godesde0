package ejercicios

import (
	"strconv"
)

func AtoiIf(valor string) (int, string) {
	num1, _ := strconv.Atoi(valor)
	var texto string
	if num1 > 100 {
		texto = "Es mayor que 100"
	} else {
		texto = "Es menor que 100"

	}
	return num1, texto
}
