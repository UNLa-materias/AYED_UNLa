package main

import "fmt"

type Nodo struct {
	dato      int
	izquierda *Nodo
	derecha   *Nodo
}

func nuevoNodo(dato int) *Nodo {
	return &Nodo{dato: dato}
}

func insertar(nodo *Nodo, dato int) {
	if dato > nodo.dato {
		if nodo.derecha == nil {
			nodo.derecha = nuevoNodo(dato)
		} else {
			insertar(nodo.derecha, dato)
		}
	} else {
		if nodo.izquierda == nil {
			nodo.izquierda = nuevoNodo(dato)
		} else {
			insertar(nodo.izquierda, dato)
		}
	}
}

func preorden(nodo *Nodo) {
	if nodo != nil {
		fmt.Printf("%d,", nodo.dato)
		preorden(nodo.izquierda)
		preorden(nodo.derecha)
	}
}

func inorden(nodo *Nodo) {
	if nodo != nil {
		inorden(nodo.izquierda)
		fmt.Printf("%d,", nodo.dato)
		inorden(nodo.derecha)
	}
}

func postorden(nodo *Nodo) {
	if nodo != nil {
		postorden(nodo.izquierda)
		postorden(nodo.derecha)
		fmt.Printf("%d,", nodo.dato)
	}
}

func buscar(nodo *Nodo, dato int) *Nodo {
	if nodo == nil {
		return nil
	}
	if nodo.dato == dato {
		return nodo
	}
	if dato > nodo.dato {
		fmt.Printf("No esta el %d != %d --> DERECHA\n", dato, nodo.dato)
		return buscar(nodo.derecha, dato)
	}
	fmt.Printf("No esta el %d != %d --> IZQUIERDA\n", dato, nodo.dato)
	return buscar(nodo.izquierda, dato)
}

