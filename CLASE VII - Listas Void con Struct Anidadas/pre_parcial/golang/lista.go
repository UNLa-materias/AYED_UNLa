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

func getPrimero[T any](lista *Lista[T]) *Nodo[T] {
	return lista.primero
}

func setPrimero[T any](lista *Lista[T], nuevoPrimero *Nodo[T]) {
	lista.primero = nuevoPrimero
}

func insertarPrimero[T any](lista *Lista[T], dato T) {
	lista.primero = crearNodo(dato, lista.primero)
}

func insertarUltimo[T any](lista *Lista[T], dato T) {
	nuevo := crearNodo(dato, nil)
	if lista.primero == nil {
		lista.primero = nuevo
		return
	}
	actual := lista.primero
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

func eliminarPrimero[T any](lista *Lista[T]) (dato T, ok bool) {
	if lista.primero == nil {
		var zero T
		return zero, false
	}
	dato = lista.primero.dato
	lista.primero = lista.primero.siguiente
	return dato, true
}

func eliminarPosicion[T any](lista *Lista[T], pos int) (dato T, ok bool) {
	tam := obtenerTamanio(lista)
	if lista.primero == nil {
		fmt.Printf("Lista vacia\n")
		var zero T
		return zero, false
	}
	if pos < 0 || pos >= tam {
		fmt.Printf("Posicion fuera de rango\n")
		var zero T
		return zero, false
	}
	if pos == 0 {
		return eliminarPrimero(lista)
	}

	anterior := (*Nodo[T])(nil)
	actual := lista.primero
	i := 0
	for i != pos {
		anterior = actual
		actual = actual.siguiente
		i++
	}
	dato = actual.dato
	anterior.siguiente = actual.siguiente
	return dato, true
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

