package ejer_interfaces

import (
	"fmt"

	"github.com/leonardprado/godesde0/interfaces"
)
func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Sou un/a %s, y estoy respirando \n", hu.Genero())
}