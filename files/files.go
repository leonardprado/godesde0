package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leonardprado/godesde0/ejercicios"
)
	var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("No se pudo crear el Archivo, Error de tipo: " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}
func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Hubo un error al concatenar el Archivo de tipo: ")
	}
}
func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error durante el Append de tipo: " + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString, de tipo: " + err.Error())
		return false
	}
	arch.Close()
	return true
}
func LeoArchivos() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error al Leer el Archivo, de tipo: " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> "+ registro)
	}
	archivo.Close()
}
