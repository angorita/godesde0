package modelos

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando   bool
	Comiendo   bool
	Vivo       bool
}

// para implementar interface humano tiene que tener las mismas funciones que humano
func (h *Hombre) Respirar()    { h.Respirando = true }
func (h *Hombre) Comer()       { h.Comiendo = true }
func (h *Hombre) Pensar()      { h.Pensando = true }
func (h *Hombre) EstaVivo()    { h.Vivo = true }
func (h *Hombre) Sexo() string { return "Hombre" }
