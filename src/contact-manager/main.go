package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()
	// PROCESO DE SERIALIZACION
	encoder := json.NewEncoder(file) // se usa para convertir estructuras de datos en formato json
	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}

	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	//PROCESO DE DESERIALIZACION
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contact

	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error al cargar los contatos: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	// mostrar opciones al usuario
	for {
		fmt.Print("======= GESTOR DE CONTACTOS ======\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")

		//Leer la respuesta del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion: ", err)
			return
		}

		//manejar la opciones del usuario
		switch option {
		case 1:
			//Ingresar y crear el contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//Agregar un contacto a Slice
			contacts = append(contacts, c)

			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2:
			// Mostramos todos los contactos
			fmt.Println("======================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Emali: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("======================")
		case 3:
			//salir del programa
			return
		default:
			fmt.Println("Opcion invalida")

		}

	}

}

func manejoErrores() {
	//MANEJO DE ERRORES

	// str := "123"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Numero: ", num)

	result, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Resultado: ", result)

	// DEFER: pospone la ejecucion de la funcion agregandola a una pila
	file, err := os.Create("hola.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() //siempre se ejecutara antes de terminar de ejecutar la funcion principal

	_, err = file.Write([]byte("Hola Ivan"))

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	/*PANIC: genera un estado de penico en el codigo cuando ocurre alguna excepcion
	RECOVER: captura el error y lo devuelve sin detener el flujo del programa
	esto solo se recomienda utilizar en casos excepcionales como manejo y limpieza de un programa*/
	divPanic(10, 2)
	divPanic(10, 0)
	divPanic(200, 5)

	// REGISTRO DE ERRORES
	log.SetPrefix("main():")                    // agrega un prefijo al log
	log.Print("Este es un mensaje de registro") // imprime el mensaje en consola

	//CREA UN ARCHIVO DE LOGS
	createLogFile()
}

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir entre cero")
	}
	return dividendo / divisor, nil
}

// OTRA FORMA DE CREAR ERRORES PERSONALIZADOS
func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("No se puede dividir entre cero, error con fmt")
	}
	return dividendo / divisor, nil
}

// PANIC
func divPanic(dividendo, divisor int) {
	// RECOVER
	defer func() { //funcion anonima
		if r := recover(); r != nil {
			fmt.Println(r) //captura el error y lo devuelve sin detener el flujo del programa
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividr entre cero")
	}
}

func createLogFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err) // es un log que detiene el flujo del programa
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("¡Oye, soy un log!")
}
