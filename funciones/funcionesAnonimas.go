package funciones

import (
	"fmt"
	//"math"
)

func Calculos() {
	/*
	suma := func (numero1 int, numero2 int) int {
		return numero1 + numero2
	}
	fmt.Print("la suma de los numeros es: ", +suma(10,25))
	*/
	var numero3 int= 32
	var numero4 int= 60
	
	calculo := func (numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println("la suma de los 4 numeros es: ", calculo(10,25))

	calculo = func(numero1, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}
	fmt.Println("el Calculo de los numeros es: ",+calculo(10, 25))
}
