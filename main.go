package main

import (
	"fmt"

	"github.com/angorita/godesde0/variables"
)
func main() {
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado,texto:=variables.ConviertoaTexto(122)
	fmt.Println(estado)
	fmt.Println(texto)

}