package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "BSAS"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   38,
		"Real Madrid": 38,
		"Chivas":      37,
	}
	fmt.Println(campeonato)

	//Estilo For Each
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]

	fmt.Printf("El puntaje Camputardo es %d, y el equipo existe = %t", puntaje, existe)
}
