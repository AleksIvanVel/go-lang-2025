package main

//importar varios paquetes
import (
	"fmt"

	//pasquete externo
	"rsc.io/quote"
)

func main() {
	fmt.Println("hola mundo")

	//funcion del paquete externo
	fmt.Println(quote.Hello())
}
