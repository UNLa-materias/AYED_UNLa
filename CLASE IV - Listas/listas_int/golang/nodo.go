package main

import "fmt"

type Nodo struct {
	dato      int
	siguiente *Nodo
}

func crearNodo(dato int, siguiente *Nodo) *Nodo {
	return &Nodo{dato: dato, siguiente: siguiente}
}

func mostrarNodo(nodo *Nodo) {
	fmt.Printf("{ %d }", nodo.dato)
}

