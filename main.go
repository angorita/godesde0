package main

import (
	"fmt"

	// "github.com/angorita/godesde0/variables"
	"github.com/angorita/godesde0/ejercicios"
)

func main() {
	/*
		variables.MostrarEnteros()
		variables.RestoVariables()
		estado, texto := variables.ConviertoaTexto(122)
		fmt.Println(estado)
		fmt.Println(texto)
		if os := runtime.GOOS; os == "Linux." || os == "OS X." {
			fmt.Println("Esto no es Windows,es ", os)
		} else {
			fmt.Println("Esto es Linux")
		}
	*/
	texto, valor := ejercicios.AtoiIf("101")
	fmt.Println(texto, valor)
}
