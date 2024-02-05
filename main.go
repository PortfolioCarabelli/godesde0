package main

// si se importa 1 paquete es asi
// import "fmt"
// si tuviera que importar varios
// import (
// 	"fmt"
// )

import (
	"fmt"
	"runtime"
)

func main() {
	// estado, text := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(text)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es linux", os)
	} else {
		fmt.Println("Es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es linux")
	case "windows":
		fmt.Println("esto es Windows")
	default:
		fmt.Printf("%s \n ", os)
	}
}
