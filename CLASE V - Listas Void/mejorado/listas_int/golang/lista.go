package main

import "fmt"

type Lista[T any] struct {
	primero *Nodo[T]
}

func crearLista[T any]() *Lista[T] {
	return &Lista[T]{primero: nil}
}

func insertarPrimero[T any](lista *Lista[T], dato T) {
	lista.primero = crearNodo(dato, lista.primero)
}

func insertarUltimo[T any](lista *Lista[T], dato T) {
	nuevo := crearNodo(dato, nil)
	actual := lista.primero
	if actual == nil {
		lista.primero = nuevo
		return
	}

	for actual.siguiente != nil {
		actual = actual.siguiente
	}
	actual.siguiente = nuevo
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

func mostrarLista[T any](lista *Lista[T], mostrar func(T)) {
	fmt.Printf("\n<LISTA> TAM:%d \n", obtenerTamanio(lista))
	actual := lista.primero
	for actual != nil {
		mostrar(actual.dato)
		actual = actual.siguiente
	}
	fmt.Printf("\n\n")
}

