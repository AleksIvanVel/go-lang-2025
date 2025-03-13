package main

import (
	"fmt"
	"log"

	greetings "github.com/AleksIvanVel/go-lang"
)

func main() {
	log.SetPrefix("greetings: ") // se utiliza para agregar un prefijo
	log.SetFlags(0)              // agrega una bandera de formato

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
