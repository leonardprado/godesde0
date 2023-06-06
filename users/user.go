package users

import (
	//"fmt"
	"fmt"
	"time"

	"github.com/leonardprado/godesde0/modelos"
)
func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Pablo", "Perez", "usuario@gmail.com", "Perez123", time.Now(), true)
	fmt.Println(u)
}
