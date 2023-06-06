package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	
	
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println( paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int {
		"Argentina": 10,
		"Mexico": 10,
		"colombia": 9,
		"Venezuela": 8,
		"Chile": 8,
		"Ecuador": 7,
		"Brasil": 6,
		"Peru": 5,
		"Bolivia": 4,
		"Canada": 3,
		"Estados Unidos": 3,
		"Uruguay": 2,
		"Paraguay": 1,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo: %s, tiene el puntaje de: %d \n", equipo, puntaje)
	}
	delete(campeonato, "Uruguay")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Costa Rica"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)

	puntaje1, existe1 := campeonato["Argentina"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje1, existe1)
}