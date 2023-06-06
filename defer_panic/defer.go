package defer_panic

import (
	"fmt"
	"log"
	//"os"+
)
func VemosDefer() {
	fmt.Println("Este es el Primer mensaje")
	defer fmt.Println("Este es el mensaje Final")
	fmt.Println("Este es el Segundo mensaje")
}
func EjemploPanic() {
	defer func ()  {
		reco := recover()
		if reco != nil {
			log.Fatalf("Error en EjemploPanic: \n %v", reco)
	}
	}()
	a:=1
	if a==1 {
		panic("Esto es un panic")
	}
}
