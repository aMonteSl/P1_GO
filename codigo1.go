package main

import (
	"fmt"
	"os"
)

// Constante tamaño del array
const SIZE = 5

// Constante para el número de opciones del menú
const OPCIONES = 7

/*
1. Implementa una función para inicializar las posiciones del array
2. Implementa una función para imprimir las posiciones del array
3. Implementa una función que devuelva la media
4. Implementa una función que devuelva el valor máximo
5. Implementa una función que devuelva el valor mínimo
6. Implementa una función que dado un valor, devuelva la posición del array
7. Escribe un programa en el que te declares un array y desde el invoques a las anteriores
funciones a modo de menú.
*/

// 1. Función para inicializar las posiciones del array
func inicializarArray(arr *[SIZE]int) {
	for i := 0; i < SIZE; i++ {
		fmt.Printf("Introduce un número para la posición %d: ", i)
		fmt.Scan(&arr[i])
	}
}

// 2. Función para imprimir las posiciones del array
func imprimirArray(arr *[SIZE]int) {
	if len(arr) == 0 {
		fmt.Println("El array está vacío")
	} else {
		for i := 0; i < SIZE; i++ {
			fmt.Printf("El valor de la posición %dº es %d\n", i, arr[i])
		}
	}

}

// 3. Función que devuelva la media
func calcularMedia(arr *[SIZE]int) float64 {
	// Definimos una variable para almacenar la suma de los valores del array
	var suma int
	// Definimos una variable para almacenar la media
	var media float64

	// Recorremos el array y sumamos los valores
	for i := 0; i < SIZE; i++ {
		suma += arr[i]
	}

	// Calculamos la media
	media = float64(suma) / float64(SIZE)

	return media

	
}

// 4. Función que devuelva el valor máximo
func calcularMaximo(arr *[SIZE]int) int{
	// Definimos una variable para almacenar el valor máximo
	var maximo int = arr[0]
	for i := 1; i < SIZE; i++ {
		if arr[i] > maximo {
			maximo = arr[i]
		}
	}

	return maximo
}

// 5. Función que devuelva el valor mínimo
func calcularMinimo(arr *[SIZE]int) int{
	// Definimos una variable para almacenar el valor mínimo
	var minimo int = arr[0]
	for i := 1; i < SIZE; i++ {
		if arr[i] < minimo {
			minimo = arr[i]
		}
	}

	return minimo
}

// 6. Función que dado un valor, devuelva la posición del array
func buscarValor(arr *[SIZE]int, valor int) int {
	// Definimos una variable para almacenar la posición del valor
	var posicion int = -1
	for i := 0; i < SIZE; i++ {
		if arr[i] == valor {
			posicion = i
			break
		}
	}

	return posicion
}

// Función para imprimir la posición del valor solicitado
func imprimirPosicion(posicion int, valor int) {
	// Si la posición es -1, significa que el valor no se encuentra en el array
	if posicion == -1 {
		fmt.Printf("El valor %d no se encuentra en el array\n", valor)
	// Si la posición es distinta de -1, significa que el valor se encuentra en la posición indicada
	} else {
		fmt.Printf("El valor %d se encuentra en la posición %d\n", valor, posicion)
	}
}


// Funcion para imprimir el valor mínimo
func imprimirMinimo(minimo int) {
	fmt.Printf("El valor mínimo es %d\n", minimo)
}

// Función para imprimir el valor máximo
func imprimirMaximo(maximo int) {
	fmt.Printf("El valor máximo es %d\n", maximo)
}

// Función para imprimir la media
func imprimirMedia(media float64) {
	fmt.Printf("La media es %.2f\n", media)
}

// Funcion que limpia la pantalla
func limpiarPantalla() {
	fmt.Println("\033[H\033[2J")
}

// Función para imprimir el menú
func imprimirMenu() {
	// Antes de imprimir el menú, limpiamos la pantalla
	limpiarPantalla()

	// Imprimimos el menú
	fmt.Println("1. Inicializar array")
	fmt.Println("2. Imprimir array")
	fmt.Println("3. Calcular media")
	fmt.Println("4. Calcular valor máximo")
	fmt.Println("5. Calcular valor mínimo")
	fmt.Println("6. Buscar valor")
	fmt.Println("7. Salir")
}

// Funcion para leer del teclado
func recivirInput(input *int) {
	// Pedimos al usuario que introduzca un input
	fmt.Printf("Introduce un valor: ")
	// Recibimos el input
	fmt.Scan(input)
}



// Funcion switch para llamar a las funciones según la opción elegida
func switchMenu(opcion int, arr *[SIZE]int) {
	// Definimos una variable para almacenar la media, maximo, minimo, valor y posicion(pide que als funciones devuelvan un valor)
	var media float64
	var maximo int
	var minimo int
	var valor int
	var posicion int


	// Limpiamos la pantalla
	limpiarPantalla()

	// Llamamos a la función correspondiente según la opción elegida
	switch opcion {
	case 1:
		// Llamamos a la función inicializarArray
		inicializarArray(arr)
	case 2:
		// Llamamos a la función imprimirArray
		imprimirArray(arr)
	case 3:
		// Llamamos a la función calcularMedia
		media = calcularMedia(arr)
		imprimirMedia(media)
	case 4:
		// Llamamos a la función calcularMaximo
		maximo = calcularMaximo(arr)
		imprimirMaximo(maximo)
	case 5:
		// Llamamos a la función calcularMinimo
		minimo = calcularMinimo(arr)
		imprimirMinimo(minimo)
	case 6:
		// Pedimos al usuario que introduzca un valor a buscar
		recivirInput(&valor)
		posicion = buscarValor(arr, valor)
		imprimirPosicion(posicion, valor)
	case 7:
		// Salimos del programa
		fmt.Println("Fin del programa")
	default:
		// Si la opción no es válida, mostramos un mensaje de error
		fmt.Println("Opción no válida")
	}
	// Esperar unos segundos para que el usuario pueda leer el mensaje
	fmt.Println("\n\nPulsa enter para continuar")
	fmt.Scanln()

}

func main() {
	// Definimos el array con el tamaño SIZE
	var array [SIZE]int
	// Definimos una variable para almacenar la opción del menú
	var opcion int

	// Mientras que la opción no sea 7, se seguirá ejecutando el programa
	for opcion != 7 {
		// Imprimimos el menú
		imprimirMenu()
		// Recivimos la opción elegida
		recivirInput(&opcion)
		// Llamamos a la función switchMenu
		switchMenu(opcion, &array)
	}

	// Salimos del programa
	os.Exit(0)

}
