package modelos

type Mujer struct{
	Hombre
	EsMadre bool

}
func(m *Mujer)Respirar(){m.respirando=true}
func(m*Mujer)Comer(){h.Comiendo=true}
func(m*Mujer)Pensar(){h.Pensando=true}
func(m*Mujer)Sexo()string{return "Mujer"}