package main

import (
	//"fmt"

	/*
		"os"
		"io"
		"bufio"
		"strings"
		"runtime"
		"github.com/leonardprado/godesde0/teclado"
		"github.com/leonardprado/godesde0/variables"
		"github.com/leonardprado/godesde0/ejercicios"
		"github.com/leonardprado/godesde0/ejercicios"
		"github.com/leonardprado/godesde0/files"
		"github.com/leonardprado/godesde0/files"
		"github.com/leonardprado/godesde0/ejercicios"
		"github.com/leonardprado/godesde0/iteraciones"
		"github.com/leonardprado/godesde0/funciones"
		"github.com/leonardprado/godesde0/arreglos_slices"
		"github.com/leonardprado/godesde0/arreglos_slices/mapas"
		"github.com/leonardprado/godesde0/users"

	*/
	e "github.com/leonardprado/godesde0/ejer_interfaces"
	"github.com/leonardprado/godesde0/modelos"
)

func main() {
	/*
		variables.MuestroEnteros()
		ejercicios.TablaMultiplicar()
		estado, texto := variables.ConviertoaTexto(1588)
		fmt.Println(estado)
		fmt.Println(texto)
		files.GrabaTabla()
		files.SumaTabla()

		 if os := runtime.GOOS; os=="Linux." || os=="OS X." {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto en Linux")
			case "darwin":
				fmt.Println("Esto en Darwin")
				default:
				fmt.Printf("%s \n", os)
		}

		numero, texto := ejercicios.DevuelveDosValores("99")
		fmt.Println(numero)
		fmt.Println(texto)

		teclado.IngresoNumeros()

		iteraciones.Iterar()

		fmt.Println(ejercicios.TablaMultiplicar())
		files.GrabaTabla()
		files.SumaTabla()
		files.LeoArchivos()
		funciones.Calculos()
		funciones.LlamarClosure()
		funciones.Exponencia(2)
		arreglos_slices.MuestroArreglos()
		arreglos_slices.MuestroSlices()
		arreglos_slices.Capacidad()
		mapas.MostrarMapas()
		users.AltaUsuario()

	*/
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
}
