package users

import (
	"fmt"
	"time"

	"github.com/angorita/godesde0/modelos"
)

//creo un objeto para usarlo

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
