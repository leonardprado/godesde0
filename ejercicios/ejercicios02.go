package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
	var numIng int
	var err error
	var texto string

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un Numero: ")
		if scanner.Scan(){
			numIng, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}else {
				break
			}
		}
	}
	for i := 1 ; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numIng, i, i*numIng)
	}
	return texto
}