package main

import (
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
		switch runtime.GOOS {
		case "linux":
			fmt.Println("Esto es linux")
		case "Darwin":
			fmt.Println("Esto es darwin")
		case "Windows":
			fmt.Println("Esto es windows")
		default:
			fmt.Println("Niguna es correcta")

		}
		texto1, valor := ejercicios.AtoiIf("11c")
		fmt.Println(texto1, valor)
		teclado.IngresoNumeros()

		iteraciones.Iterar()
		iteraciones.IterarBreak()
		iteraciones.IterarContinue()
	*/
	ejercicios.Tablas(6)
	ejercicios.Ingreso()
}
