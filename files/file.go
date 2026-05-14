package files

import (
	"bufio"
	"fmt"
	"godesde0-practica/ejercicios"
	"os"
)

var fileUrl string = "./files/txt/tabla.txt"

func GrabaTabla() {
	texto := ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileUrl)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTablaArchivo() {
	texto := ejercicios.TablaMultiplicar()
	if !Append(fileUrl, texto) {
		fmt.Println("Error al agregar el texto al archivo")
	}
}

func Append(filen string, texto string) bool {
	archivo, err := os.OpenFile(fileUrl, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	archivo.Close()

	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileUrl)

	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
		archivo.Close()
	}
}
