package defer_panic

import (
	"fmt"
	"log"
)

//Defer: Instruccion que nos permite decir cual es la ultima instruccion al ejecutar

func VemosDefer() {

	fmt.Println("este es un primer Mensaj")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo")

}

// Panic: Aborta el programa con un mensaje en cosola
// Recover: me permite recuperarme de un panic
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Hubo un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Hola")
	}
}
