package funciones

//Que son?
// son funciones que no tienen un nombre que pueden ser asignada tanto a una variable y pueden pasarse por parametro

import "fmt"

func Calculos(){
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(10,10))
}
