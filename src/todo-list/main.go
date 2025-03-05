package main

import (
	"bufio"
	"fmt"
	"os"
)

// ESTRUCTURAS
type persona struct {
	nombre string
	edad   string
	correo string
}

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// METODOS: son propios de las estructuras
func (p *persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

// metodo para agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para marcar como completado
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	bases()

	//intancia de la lista de tareas
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("Selecciones una opcion:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea com completava\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")

		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")

		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea editar: ")
			fmt.Scanln(&index)

			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')

			lista.editarTarea(index, t)
			fmt.Println("Tarea editada correctamente")

		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)

			lista.eliminarTarea(index)

		case 5:
			fmt.Print("Saliendo del programa...")
			return

		default:
			fmt.Print("Opcion invalida")

		}

		fmt.Println("Lista de tareas")
		fmt.Println("================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d-. %s - %s - completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("================================")
	}
}

func bases() {
	//declaracion de matrices
	var a = [5]int{10, 20, 30, 40, 50}

	//matriz bidimencional
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)

	//accedemos al indice y al valor del elemento
	for index, value := range a {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

	// slices
	diasSemana := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}

	// sub slice de diasSemana
	finSemana := diasSemana[5:7]

	//agregar elemento a la slice
	finSemana = append(finSemana, "Viernes")

	fmt.Println(finSemana)

	//imprime tamanio de la slice
	fmt.Println("Tamanio de la slice: ", len(finSemana))

	//imprime la capacidad maxima de la slice
	fmt.Println("Capacidad total de la slice: ", cap(finSemana))

	// MAPAS
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	fmt.Println(colors["rojo"])

	// agregar elemento
	colors["negro"] = "#000000"

	// eliminar elemento
	delete(colors, "rojo")

	fmt.Println(colors)

	// ESTRUCTURAS
	p := persona{"aleks", "21", "aleks@example.com"}
	fmt.Println(p)
	p.diHola()

	// PUNTEROS: acceder a una varible mediante su referencia en memoria
	var x int = 10
	fmt.Println(x)
	editar(&x) //modifica el valor de la var mediante su referencia en memoria
	fmt.Println(x)
}

func editar(x *int) {
	*x = 20
}
