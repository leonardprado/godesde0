package variables

import (
	"fmt"
	"time"
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Sueldo64 float64

var Fecha time.Time


func RestoVariables() {
	Nombre = "Leonardo"
	Estado = true
	Sueldo = 20000
	Fecha = time.Now()
	Sueldo64 = 200000

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return true, texto
}