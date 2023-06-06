package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2,5,4}
var arreglo [10]int = [10]int{6,98,21,674,15,36,78,9}

func MuestroSlices() {
	fmt.Println(tablaSlice)

	porcion1 := arreglo[3:]	//Slice creado desde un vector, desde la posicion 3
	porcion2 := arreglo[:6] //Slice creado desde un vector, desde la posicion 0 hasta la 6
	porcion3 := arreglo[3:8] //Slice creado desde un vector, desde la posicion 3 hasta la 8

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}
func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("\nLargo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 2, 2)
	for i := 0; i < 200; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo Slice %d, Capacidad %d", len(nums), cap(nums))
}
