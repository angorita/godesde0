package funciones

import "fmt"

//devuelve una funcion anonima que devuelve un valor
func tabla(valor float64)func()float64{
	//esto queda oculto 
	numero:=valor
	secuencia:=0
	return func() float64 {
		secuencia++
		return numero*float64(secuencia)
	}
	
}
func LlamarClosure(){
	tabladel:=2
	mitabla:=tabla(float64(tabladel))
	for i:=1;i<=10;i++{
		fmt.Println(mitabla())
	}
}