package modelos

//va a tener las mismas propiedades del hombre

type Mujer struct{
	Hombre //hereda propiedades
	EsMadre bool //propiedad propia 
	TieneConcha bool

}
func(m *Mujer)Respirar(){m.Respirando=true}
func(m *Mujer)Comer(){m.Comiendo=true}
func(m *Mujer)Pensar(){m.Pensando=true}
func(m *Mujer)Vive(){m.Vivo=true}
func(m *Mujer)Sexo()string{return "Femenino"}