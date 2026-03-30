package main

import "fmt"

type ArbolAVL struct {
	dato      int
	izquierdo *ArbolAVL
	derecho   *ArbolAVL
}

func insertar(unArbol **ArbolAVL, dato int) bool {
	if *unArbol == nil {
		*unArbol = &ArbolAVL{dato: dato}
	} else {
		if (*unArbol).dato < dato {
			insertar(&((*unArbol).derecho), dato)
		} else {
			insertar(&((*unArbol).izquierdo), dato)
		}
	}

	balancearArbolAVL(unArbol)
	return true
}

func buscar(unArbol *ArbolAVL, dato int) bool {
	if unArbol == nil {
		return false
	}
	if unArbol.dato == dato {
		return true
	}
	if unArbol.dato < dato {
		return buscar(unArbol.derecho, dato)
	}
	return buscar(unArbol.izquierdo, dato)
}

func mostrar(unArbol *ArbolAVL) {
	if unArbol == nil {
		return
	}
	fmt.Printf(" \t%d ", unArbol.dato)
	mostrar(unArbol.derecho)
	mostrar(unArbol.izquierdo)
}

func alturaArbolAVL(unArbol *ArbolAVL) int {
	if unArbol == nil {
		return 0
	}
	alturaIzq := alturaArbolAVL(unArbol.izquierdo)
	alturaDer := alturaArbolAVL(unArbol.derecho)
	if alturaDer > alturaIzq {
		return alturaDer + 1
	}
	return alturaIzq + 1
}

func balanceo(unArbol *ArbolAVL) int {
	if unArbol == nil {
		return 0
	}
	return alturaArbolAVL(unArbol.derecho) - alturaArbolAVL(unArbol.izquierdo)
}

func balancearArbolAVL(unArbol **ArbolAVL) {
	if unArbol == nil || *unArbol == nil {
		return
	}

	auxBalance := balanceo(*unArbol)
	if auxBalance > 1 {
		if balanceo((*unArbol).derecho) >= 1 {
			rotarIzq(unArbol)
		} else {
			rotarDerecha(&((*unArbol).derecho))
			rotarIzq(unArbol)
		}
	} else if auxBalance < -1 {
		if balanceo((*unArbol).izquierdo) <= -1 {
			rotarDerecha(unArbol)
		} else {
			rotarIzq(&((*unArbol).izquierdo))
			rotarDerecha(unArbol)
		}
	}
}

func rotarDerecha(unArbol **ArbolAVL) bool {
	if unArbol == nil || *unArbol == nil || (*unArbol).izquierdo == nil {
		return false
	}
	raiz := *unArbol
	raizNueva := raiz.izquierdo
	auxiliar := raizNueva.derecho

	*unArbol = raizNueva
	raizNueva.derecho = raiz
	raiz.izquierdo = auxiliar
	return true
}

func rotarIzq(unArbol **ArbolAVL) bool {
	if unArbol == nil || *unArbol == nil || (*unArbol).derecho == nil {
		return false
	}
	raiz := *unArbol
	raizNueva := raiz.derecho
	auxiliar := raizNueva.izquierdo

	*unArbol = raizNueva
	raizNueva.izquierdo = raiz
	raiz.derecho = auxiliar
	return true
}

