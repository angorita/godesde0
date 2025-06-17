package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	//instruccion que permite ultima ejecucion y cerrar
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el msg final")

	fmt.Println("Este es el segundo mensaje")
}
func EjemploPanic(){
	defer func(){
		reco:=recover()
		//el recover siempre se usa con defer
		if reco!=nil{
			log.Fatalf("ocurrio un error que genero un panic\n%v",reco)
		}
	}()
	a:=1
	if a==1{panic("Se encontro uno,aborto el mensaje")}
}

