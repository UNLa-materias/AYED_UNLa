package main

import "fmt"

type Nodo struct {
	dato      int
	siguiente *Nodo
}

func crearNodo(dato int, siguiente *Nodo) *Nodo {
	return &Nodo{dato: dato, siguiente: siguiente}
}

func getDato(n *Nodo) int {
	return n.dato
}

func getSiguiente(n *Nodo) *Nodo {
	return n.siguiente
}

func setDato(n *Nodo, nuevoDato int) {
	n.dato = nuevoDato
}

func setSiguiente(n *Nodo, nuevoSig *Nodo) {
	n.siguiente = nuevoSig
}

func mostrarNodo(n *Nodo) {
	fmt.Printf("{ %d }", n.dato)
}

