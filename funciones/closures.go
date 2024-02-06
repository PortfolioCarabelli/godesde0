package funciones

import "fmt"

//Devuleve una funcion y esa funcion devuelve un entero

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {

	tabladel := 2

	// Asigne a una funcion a una variable
	MiTabla := tabla(tabladel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
