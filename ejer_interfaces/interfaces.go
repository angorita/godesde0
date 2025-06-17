package ejer_interfaces

import (
	"fmt"

	"github.com/angorita/godesde0/interfaces"
)

//las interfaces son modelos de comportamiento de una estructura
//interactuar
func HumanosRespirando(hu interfaces.Humano){
	hu.Respirar()
	fmt.Printf("Soy un %s,y estoy respirando \n",hu.Sexo(),)	

}