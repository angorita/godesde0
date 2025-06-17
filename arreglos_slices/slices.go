package arreglosslices

import "fmt"

var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestoSlice() {
	porcion := arreglo[3:]   //creado desde un vector desde la posicion 3
	porcion2 := arreglo[:5]  //desde la posicion 0 a la 5
	porcion3 := arreglo[6:7] //slice desde 6 hasta 7
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}
