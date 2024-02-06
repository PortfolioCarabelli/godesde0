package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 100000000 {
		return
	}

	Exponencia(valor * 2)
	fmt.Println(valor)
}
