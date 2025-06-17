package arreglosslices

import "fmt"

//slice no tiene dimension vecrores
var tabla[10]int=[10]int{10,0,59}
var matriz[5][10]string
func MuestroArreglos(){
	tabla[7]=33
	tabla[2]=54//pisa el 59
	fmt.Println(tabla)
	tabla2:=[10]string{"Mengeche","Os","Carla"}// resto cadena vacia
	fmt.Println(tabla2)
for i:=0;i<len(tabla);i++{
	fmt.Println(tabla[i])
}
matriz[2][9]="Ene"
matriz[0][9]="Feb"
matriz[4][9]="Mar"
matriz[0][0]="Abril"
matriz[0][1]="Mayo"
matriz[0][2]="Junio"
fmt.Println(matriz,"\n",matriz)
}