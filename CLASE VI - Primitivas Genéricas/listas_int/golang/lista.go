package main

import "fmt"

type Nodo[T any] struct {
	dato      T
	siguiente *Nodo[T]
}

func crearNodo[T any](dato T, siguiente *Nodo[T]) *Nodo[T] {
	return &Nodo[T]{dato: dato, siguiente: siguiente}
}

type Lista[T any] struct {
	primero *Nodo[T]
}

func crearLista[T any]() *Lista[T] {
	return &Lista[T]{primero: nil}
}

func insertarPrimero[T any](lista *Lista[T], dato T) {
	lista.primero = crearNodo(dato, lista.primero)
}

func obtenerTamanio[T any](lista *Lista[T]) int {
	tam := 0
	actual := lista.primero
	for actual != nil {
		tam++
		actual = actual.siguiente
	}
	return tam
}

func mostrarListaGenerica[T any](lista *Lista[T], mostrar func(T)) {
	fmt.Printf("\n<LISTA> TAM:%d \n", obtenerTamanio(lista))
	actual := lista.primero
	for actual != nil {
		mostrar(actual.dato)
		actual = actual.siguiente
	}
	fmt.Printf("\n\n")
}

func ordenarLista[T any](lista *Lista[T], mayor func(a, b T) bool) {
	if obtenerTamanio(lista) < 2 {
		return
	}

	var ultimo *Nodo[T]
	for {
		permutacion := false
		nodo := lista.primero
		for nodo != nil && nodo.siguiente != ultimo {
			if nodo.siguiente == nil {
				break
			}
			if mayor(nodo.dato, nodo.siguiente.dato) {
				aux := nodo.dato
				nodo.dato = nodo.siguiente.dato
				nodo.siguiente.dato = aux
				permutacion = true
			}
			nodo = nodo.siguiente
		}
		ultimo = nodo
		if !permutacion {
			break
		}
	}
}

