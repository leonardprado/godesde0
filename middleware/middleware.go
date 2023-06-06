package middleware

import "fmt"

func sumar(a,b int) int {
	return a + b
}
func restar(a,b int) int {
	return a - b
}
func multiplicar(a,b int) int {
	return a * b
}
func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMidd(sumar)(2,3)
	fmt.Println(result)
	result = operacionesMidd(sumar)(10,6)
	fmt.Println(result)
	result = operacionesMidd(sumar)(3,4)
	fmt.Println(result)
}
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a,b)
	}
} 