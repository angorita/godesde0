package main

import (
	"fmt"
	"runtime"
)

func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	// estado,texto:=variables.ConviertoaTexto(122)
	// fmt.Println(estado)
	// fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows es ", os)
	} else {
		fmt.Println("Esto es Linux")
	}
}
