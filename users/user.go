package users

import (
	"time"

	"github.com/angorita/godesde0/modelos"
)

func AltaUsuario() {
u:=new(modelos.User)
u.AddUser(10,"Pablo",time.Now(),true)
}
