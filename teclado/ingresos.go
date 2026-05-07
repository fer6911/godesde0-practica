package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresarNumeros() {
	fmt.Println("Ingrese numero 1:")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Dato ingresado incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2:")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Dato ingresado incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese una leyenda:")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
