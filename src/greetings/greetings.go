package greetings

import (
	"errors"
	"fmt"
)

// devuelve un saludo para una persona especifica
func Hello(name string) (string, error) {

	//Manejo de error en caso de enviar un nombre vacio
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
	return message, nil
}
