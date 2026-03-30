package main

import "fmt"

type Lista struct {
	primero *Nodo
}

func crearLista() *Lista {
	return &Lista{primero: nil}
}

func insertarPrimero(lista *Lista, dato int) {
	nuevo := crearNodo(dato, lista.primero)
	lista.primero = nuevo
}

func mostrarLista(lista *Lista) {
	fmt.Printf("\n<LISTA>\n")
	actual := lista.primero
	for actual != nil {
		mostrarNodo(actual)
		actual = actual.siguiente
	}
	fmt.Printf("\n\n")
}

