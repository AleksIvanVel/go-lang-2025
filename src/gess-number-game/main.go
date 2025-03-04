package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// estructuras if - else
	// if t := time.Now(); t.Hour() < 12 {
	// 	fmt.Println("Mañana")
	// } else if t.Hour() < 17 {
	// 	fmt.Println("Tarde")
	// } else {
	// 	fmt.Println("Noche")
	// }

	// estructura switch
	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Go run --> Windows")
	// case "linux":
	// 	fmt.Println("Go run --> Linux")
	// case "darwin":
	// 	fmt.Println("Go run --> macOS")
	// default:
	// 	fmt.Println("Go run --> Otro OS")
	// }

	// bulce for
	// bucle infinito
	// for{

	// }

	//for con una condicion
	// var i int
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//for convencional
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// fmt.Print(hello("Shell"))
	// sum, mul := calc(4, 5)

	// fmt.Println("la suma es: ", sum, "\nla multiplicacions es: ", mul)

	jugar()

}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero ingresado es menor")
		}
	}

	fmt.Println("Se acabaron los intentos. El número era: ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string

	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Eleccion invalida. Intente nuevamente")
		jugarNuevamente()
	}
}

// Declaracion de funciones
// func hello(name string) string {
// 	return "Hola " + name
// }

// funcion multi retorno
// podemos devolver dos variables como tal
// func calc(a, b int) (sum, mul int) {
// 	sum = a + b
// 	mul = a * b
// 	return
// }
