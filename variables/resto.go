package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Variables globales de solo resto
var Nombre string
var Estado bool
var Sueldo float32

// Para fechas y horas importar un paquete
var Fecha time.Time

func RestoVariables() {
	Nombre = "Fede"
	Estado = true
	Sueldo = 1566.66
	Fecha = time.Now()

	fmt.Println("Mi nombre es: ", Nombre)
	fmt.Println("Mi estado es: ", Estado)
	fmt.Println("Mi sueldo es: ", Sueldo)
	fmt.Println("La fecha es: ", Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string

	texto = strconv.Itoa(numero)

	return true, texto
}
