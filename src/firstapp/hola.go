package main

//importar varios paquetes
import (
	"fmt"

	//pasquete externo
	"rsc.io/quote"
)

// declaracion e variables
var nombre, apellido string
var edad int

func main() {
	fmt.Println("hola mundo")

	//funcion del paquete externo
	fmt.Println(quote.Hello())

	//otra forma de declarar e inicializar variables en una sola linea
	var (
		carrera  = " ICO "
		semestre = " octavo "
		promedio = 9.64
	)

	//otra forma de declarar e inicializar variables en una sola linea
	/*
		la palabra reservada "var" perimite declara variables dentro y fuera de las funciones
	*/
	var mascota = " lucrecia "

	//otra forma de declarar e inicializar variables en una sola linea
	/*
		la sintaxis := solo permite declara variables dentro de funciones
	*/
	vehiculo := " RC200 "

	nombre = "Aleks"
	apellido = " Velazquez "
	edad = 21

	fmt.Print(nombre, apellido, edad, carrera, semestre, promedio, mascota, vehiculo)
}
