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

func obtenerTamanio(lista *Lista) int {
	i := 0
	actual := lista.primero
	for actual != nil {
		i++
		actual = actual.siguiente
	}
	return i
}

func insertarUltimo(lista *Lista, dato int) {
	if lista.primero == nil {
		insertarPrimero(lista, dato)
		return
	}

	actual := lista.primero
	for actual.siguiente != nil {
		actual = actual.siguiente
	}
	actual.siguiente = crearNodo(dato, nil)
}

func insertarPosicion(lista *Lista, dato int, posicion int) {
	tam := obtenerTamanio(lista)

	if posicion > tam+1 {
		fmt.Printf("\n\nERROR, no se puede insertar en pos que no existen!!!")
		return
	}
	if posicion < 0 {
		fmt.Printf("\n\nERROR, no se puede insertar en pos negativas!!!")
		return
	}

	if posicion == 0 {
		insertarPrimero(lista, dato)
		return
	}
	if posicion == tam {
		insertarUltimo(lista, dato)
		return
	}

	if posicion > 0 && posicion < tam {
		actual := lista.primero
		for i := 0; i < posicion-1; i++ {
			actual = actual.siguiente
		}
		actual.siguiente = crearNodo(dato, actual.siguiente)
	}
}

func ordenarLista(lista *Lista) {
	tam := obtenerTamanio(lista)
	if tam < 2 {
		return
	}

	for i := 0; i < tam; i++ {
		actual := lista.primero
		for j := 0; j < tam-1; j++ {
			if actual != nil && actual.siguiente != nil && actual.dato > actual.siguiente.dato {
				aux := actual.dato
				actual.dato = actual.siguiente.dato
				actual.siguiente.dato = aux
			}
			actual = actual.siguiente
		}
	}
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

