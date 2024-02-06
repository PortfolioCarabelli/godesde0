package main

import "github.com/PortfolioCarabelli/godesde0/middleware"

// si se importa 1 paquete es asi
// import "fmt"
// si tuviera que importar varios
// import (
// 	"fmt"
// )

func main() {
	// estado, text := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(text)

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("No es linux", os)
	// } else {
	// 	fmt.Println("Es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("esto es linux")
	// case "windows":
	// 	fmt.Println("esto es Windows")
	// default:
	// 	fmt.Printf("%s \n ", os)
	// }

	// num, text := ejercicios.ConvertiraEntero("1000")
	// fmt.Println(num)
	// fmt.Println(text)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// files.LeoArchivoOs()

	// Pedro := new(models.Hombre)

	// ejer_interfaces.HumanosRespirando(Pedro)

	// Maria := new (models.Mujer)

	// ejer_interfaces.HumanosRespirando(Maria)

	// defer_panic.EjemploPanic()
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentoo("Federico Carabelli", canal1)
	// defer func() {
	// 	// esto es un await
	// 	<-canal1
	// }()
	middleware.MiMiddleware()
}
