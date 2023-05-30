package ejercicios


import "strconv"

func DevuelveDosValores(texto string) (int, string) {
num, err := strconv.Atoi(texto)
if err != nil {
	return 0, "Hubo un error " + err.Error()
}
if num > 100 {
	return num, "Es mayor a 100"
} else if (num == 100) {
	return num, "Es igual a 100"
} else {
	return num, "Es menor a 100"
}
}





