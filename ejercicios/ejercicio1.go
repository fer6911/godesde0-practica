package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio1(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)

	fmt.Println(texto, "el valor del parametro es:", numero)
	if err != nil {
		return 0, "Error: No se pudo convertir el texto a número"
	}
	if numero > 100 {
		return numero, "El número es mayor que 100"
	}
	return numero, "El número es menor o igual a 100"
}
