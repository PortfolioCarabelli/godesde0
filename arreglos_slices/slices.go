package arreglos_slices

import "fmt"

// son vectores que no tienen una capacidad incial

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 10, 6, 20, 50}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3 hasta el final
	procion2 := arreglo[:5]  //Slice creado desde un vector, desde la posicion 0 a 5
	procion3 := arreglo[6:7] //Slice creado desde un vector, desde la posicion 6 a 7
	fmt.Println(porcion)
	fmt.Println(procion2)
	fmt.Println(procion3)
}

func Capacidad() {
	// va a tener 5 elementos y va a tener de capacidad de 20
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n Largo %d, Capacidad%d", len(nums), cap(nums))
}
