package main

import (
	"github.com/angorita/godesde0/webserver"
)

func main() {
	/*
		variables.MostrarEnteros()
		variables.RestoVariables()
		estado, texto := variables.ConviertoaTexto(122)
		fmt.Println(estado)
		fmt.Println(texto)
		if os := runtime.GOOS; os == "Linux." || os == "OS X." {
			fmt.Println("Esto no es Windows,es ", os)
		} else {
			fmt.Println("Esto es Linux")
		}
		switch runtime.GOOS {
		case "linux":
			fmt.Println("Esto es linux")
		case "Darwin":
			fmt.Println("Esto es darwin")
		case "Windows":
			fmt.Println("Esto es windows")
		default:
			fmt.Println("Niguna es correcta")

		}
		texto1, valor := ejercicios.AtoiIf("11c")
		fmt.Println(texto1, valor)
		teclado.IngresoNumeros()

		iteraciones.Iterar()
		iteraciones.IterarBreak()
		iteraciones.IterarContinue()
		ejercicios.Tablas(6)
		ejercicios.Ingreso()
		funciones.Calculos()
		funciones.LlamarClosure()
		funciones.Exponencia(2)
		arreglosslices.MuestroArreglos()
		arreglosslices.MuestoSlice() //ajusta automaticamente...
		arreglosslices.MostrarMapas()
		users.AltaUsuario()
		Pedro := new(modelos.Hombre)
		e.HumanosRespirando(Pedro)
		Gabita := new(modelos.Mujer)
		e.HumanosRespirando(Gabita)
		e.VemosDefer()
		e.EjemploPanic()
		canal1 := make(chan bool)
		go goroutines.MiNombreLentoo("Oscar", canal1)
		// var x string
		// fmt.Scanln(&x)//cuando aprete enter se cierra excepto por los canales
		defer func(){<-canal1}() //ahora es el canal el que envia informacion
		//usando defer impide que haya algun canal abierto

		fmt.Println("Estoy aqui",canal1)
	*/
	webserver.MiWebServer()
}
