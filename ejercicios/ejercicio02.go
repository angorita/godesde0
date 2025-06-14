package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// func que reciba un int, valide si error o no y que vuelva a pedirlo
// crear la tabla numerica del mismo de 1 a 10 y la muestre por pantalla
// en main la llamo solicitar leer proceso imprimir hay llamados a otras funciones
func Tablas(tabla int) (int, error) {
	fmt.Println("Esta es la tabla del : ", tabla)
	for i := 1; i < 11; i += 1 {
		fmt.Println(tabla, i, i*tabla)
	}
	return tabla, nil
}
func Ingreso(){
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese algo")
	if scanner.Scan(){
		num1,err:=strconv.Atoi(scanner.Text())
		if err!=nil{
			panic("Error en los ..."+err.Error())
		}
	Tablas(num1)
		
	}


}
