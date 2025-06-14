package ejercicios

import (
	"strconv"
)

func AtoiIf(valor string) (int, string) {
	num1, e := strconv.Atoi(valor)
	if e != nil {
		return 0, "Hubo un error importante fijate!!! --> " + e.Error()
	}
	var texto string
	if num1 > 100 {
		texto = "Es mayor que 100"
	} else {
		texto = "Es menor que 100"

	}
	return num1, texto
}
