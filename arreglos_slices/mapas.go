package arreglosslices

import "fmt"

func MostrarMapas(){
	//se puede definir con make
	//o simplemente instanciandolo
	//clave valor
	//muestra el valor
	paises:=make(map[string]string)
	paises["argentina"]="General Acha"
	paises["mexico"]="Mexico DF"
	paises["dominicana"]="Santo Domingo"
	paises["peru"]="Lima"
	paises["chile"]="Santiago"
	paises["zambia"]="Samudia"
	fmt.Print(paises["argentina"],"\n")
	fmt.Println(paises)
	campeonato:=map[string]string{
		"river":"amadeo carrizo",
		"boca":"schelotto",
		"chacarita":"no se",
	}
	fmt.Println(campeonato)
	for c,v:=range campeonato{
		fmt.Printf("El campeonato %v lo gana %s",c,v)
	}
	delete(campeonato,"river")
	
}