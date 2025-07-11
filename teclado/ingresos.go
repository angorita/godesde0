package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1, num2 int
var leyenda string
var err error

// metodo cuando es vacio
func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Hubo un error\n" + err.Error())
		}
	}
	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Hubo un error\n" + err.Error())
		}
	}
	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
		if err != nil {
			panic("Hubo un error\n" + err.Error())
		}
	}
	fmt.Println(leyenda, num1*num2)

}
