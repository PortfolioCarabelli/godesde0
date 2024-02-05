package main

// si se importa 1 paquete es asi
// import "fmt"
// si tuviera que importar varios
// import (
// 	"fmt"
// )

import (
	"fmt"

	"github.com/PortfolioCarabelli/godesde0/variables"
)

func main() {
	estado, text := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(text)
}
