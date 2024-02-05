package variables

import(
	"fmt"
)

func MostrarEnteros(){
	//Declaracion
	var intComun int
	
	//Asignacion
	intde32 := int32(10)
	intde64 := int64(100)

	//Display por consola
	fmt.Println("int comun =", intComun)
	fmt.Println("int 32 =", intde32)
	fmt.Println("int 64 =", intde64)
}