package interfaces
//comienza como una estructura
//a diferencia no colocamos variables solo definicion de funciones
type Humano interface{
	//solo definicion de funciones
	Respirar()
	Pensar()
	Comer()
	Sexo ()string
	EstaVivo()bool
}
