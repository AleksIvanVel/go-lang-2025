package main

//importar varios paquetes
import (
	"fmt"
	"math"
	"strconv"

	//paquete externo
	"rsc.io/quote"
)

// declaracion e variables
var nombre, apellido string
var edad int

// declaracion de constantes
const Pi = 3.14 //nombre del const empieza con mayus

// declaracion de constantes incremntales
const (
	Domingo = iota + 1
	Lunes
	MArtes
	Miercoles
	Jueves
	Viernes
	Sabado
)

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

	fmt.Println(nombre, apellido, edad, carrera, semestre, promedio, mascota, vehiculo, Viernes)

	//guarda el valor ascii
	var a byte = 'a'
	fmt.Println("valor ascci de a: ", a)

	//accede al primer caracter del string en su valor ascii
	s := "string"
	fmt.Println("valor ascii de un elemnto de un stingr: ", s[0])

	// guarda valor UNICODE
	var c rune = '♡'
	fmt.Println("Valor unicode de ♡: ", c)

	//valores predeterminados
	var (
		defaultInt    int
		defaultUint   uint //numeros enteros positivos
		defaulFloat   float32
		defaulBool    bool
		defaultString string
	)
	fmt.Println(defaultInt, defaultUint, defaulFloat, defaulBool, defaultString)

	//conversion de datos
	var integer16 int16 = 50
	var integer32 int32 = 100
	fmt.Println(int32(integer16) + integer32)

	//convertir un string a un entero
	numberstr := "100"
	i, _ := strconv.Atoi(numberstr)
	fmt.Println(i + 1)

	//convertir un entero a string
	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s) //concatenacion

	// formato de salida en consola con paquete fmt
	fmt.Printf("Hola me llamo %s y tengo %d años \n", nombre, edad)

	// ingresar datos por consola
	var name string
	var age int

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&name)

	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola me llamo %s y tengo %d años \n", name, age)

	// ejercicio numero 1

	var ca, co, hip, area, perimetro float64

	fmt.Println("Ingrese la longitud del cateto adyacente: ")
	fmt.Scan(&ca)

	fmt.Println("Ingrese la longitud del cateto opuesto: ")
	fmt.Scan(&co)

	hip = math.Sqrt(math.Pow(ca, 2) + math.Pow(co, 2))
	area = (ca * co) / 2
	perimetro = ca + co + hip

	fmt.Printf("La longitud de la hipotenusa es: %.2f \n", hip)
	fmt.Printf("El area del triangulo es: %.2f \n", area)
	fmt.Printf("El periemtro del triangulo es: %.2f \n", perimetro)

}
