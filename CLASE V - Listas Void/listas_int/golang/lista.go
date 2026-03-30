package main

import "fmt"

type Lista struct {
	primero *Nodo
}

func crearLista() *Lista {
	return &Lista{primero: nil}
}

func insertarPrimero(lista *Lista, dato int) {
	lista.primero = crearNodo(dato, lista.primero)
}

func insertarUltimo(lista *Lista, dato int) {
	nuevo := crearNodo(dato, nil)
	actual := lista.primero
	if actual == nil {
		lista.primero = nuevo
		return
	}

	for getSiguiente(actual) != nil {
		actual = getSiguiente(actual)
	}
	setSiguiente(actual, nuevo)
}

func obtenerTamanio(lista *Lista) int {
	tam := 0
	actual := lista.primero
	for actual != nil {
		tam++
		actual = getSiguiente(actual)
	}
	return tam
}

func ordenarLista(lista *Lista) {
	if obtenerTamanio(lista) < 2 {
		return
	}

	var ultimo *Nodo
	for {
		permutacion := false
		nodo := lista.primero
		for getSiguiente(nodo) != ultimo {
			if getDato(nodo) > getDato(getSiguiente(nodo)) {
				aux := getDato(nodo)
				setDato(nodo, getDato(getSiguiente(nodo)))
				setDato(getSiguiente(nodo), aux)
				permutacion = true
			}
			nodo = getSiguiente(nodo)
		}
		ultimo = nodo
		if !permutacion {
			break
		}
	}
}

func mostrarLista(lista *Lista) {
	fmt.Printf("\n<LISTA> TAM:%d \n", obtenerTamanio(lista))
	actual := lista.primero
	for actual != nil {
		mostrarNodo(actual)
		actual = getSiguiente(actual)
	}
	fmt.Printf("\n\n")
}

