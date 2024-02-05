package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TabladeMultiplicar() {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese n° 1: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic("error: " + err.Error())
			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d \n", numero, i, i*numero)
	}

}

func TabladeMultiplicarenTexto() string {
	var numero int
	var err error
	var text string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese n° 1: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic("error: " + err.Error())
			} else {
				break
			}
		}
	}

	for i := 1; i < 10; i++ {
		text += fmt.Sprintf("%d * %d = %d \n", numero, i, i*numero)

	}
	return text
}
