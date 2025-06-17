package funciones

import "fmt"

func Exponencia(valor float64){
	if valor>10000000000000{
		return 
	}
	fmt.Println("Oscar Angarita",valor)
	Exponencia(valor*2)
}