package main

import "fmt"

func main() {
	lista := crearLista[int]()
	lista2 := crearLista[string]()
	lista3 := crearLista[float32]()

	mostrarLista(lista, func(v int) { fmt.Printf("{ %d }", v) })
	insertarPrimero(lista, 4)
	insertarPrimero(lista, 11)
	insertarUltimo(lista, 109)
	insertarPrimero(lista, 21)
	mostrarLista(lista, func(v int) { fmt.Printf("{ %d }", v) })

	mostrarLista(lista2, func(s string) { fmt.Printf("{ %s }", s) })
	insertarPrimero(lista2, "nico")
	insertarPrimero(lista2, "brenda")
	insertarUltimo(lista2, "walter")
	insertarPrimero(lista2, "flor")
	mostrarLista(lista2, func(s string) { fmt.Printf("{ %s }", s) })

	mostrarLista(lista3, func(f float32) { fmt.Printf("{ %.2f }", f) })
	insertarPrimero(lista3, 4.12)
	insertarPrimero(lista3, 11.43)
	insertarUltimo(lista3, 109.009)
	insertarPrimero(lista3, 21.14)
	mostrarLista(lista3, func(f float32) { fmt.Printf("{ %.2f }", f) })
}

