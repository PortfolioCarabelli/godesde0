package ejercicios

import (
	"strconv"
)

func ConvertiraEntero(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	// Manejo de errores
	if err != nil {
		return 0, "Hubo un error"
	}
	//Verifico el resultado de la conversion
	if num > 100 {
		return num, "Mayor a 100"
	} else {
		return num, "menor a 100"
	}
}
