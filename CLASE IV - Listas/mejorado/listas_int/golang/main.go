package main

import "fmt"

func main() {
	lista := crearLista()

	fmt.Printf("\nTAM: %d", obtenerTamanio(lista))
	mostrarLista(lista)

	insertarUltimo(lista, 34)
	insertarPrimero(lista, 4)
	insertarPrimero(lista, 1)
	insertarUltimo(lista, 11)
	insertarPrimero(lista, -26)
	insertarUltimo(lista, 64)
	insertarPrimero(lista, 96)

	fmt.Printf("\nTAM: %d", obtenerTamanio(lista))
	mostrarLista(lista)

	insertarPosicion(lista, 44, 2)
	insertarPosicion(lista, 85, 0)
	insertarPosicion(lista, 101, 9)
	insertarPosicion(lista, 85, -1)
	insertarPosicion(lista, 85, 12)
	mostrarLista(lista)

	ordenarLista(lista)
	mostrarLista(lista)
}

