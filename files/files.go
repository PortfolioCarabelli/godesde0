package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	ejercicios "github.com/PortfolioCarabelli/godesde0/Ejercicios"
)

var fileName string = "./files/archivostxt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.TabladeMultiplicarenTexto()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TabladeMultiplicarenTexto()
	if Append(fileName, texto) == false {
		fmt.Println("error al concatenar contenido")
	}
}

func Append(filen string, text string) bool {
	//Abrimos el archivo con todos los permisos en modo escritura
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}
	//Escribo el archivo
	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error durante el write" + err.Error())
		return false
	}
	//Cierro el archivo
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivoOs() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
}
