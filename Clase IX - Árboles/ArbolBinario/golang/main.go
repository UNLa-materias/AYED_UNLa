package main

import "fmt"

func main() {
	fmt.Printf("----- Ejemplo arbol binario ----\n")

	raiz := nuevoNodo(28)
	insertar(raiz, 11)
	insertar(raiz, 96)
	insertar(raiz, 21)
	insertar(raiz, 6)
	insertar(raiz, 97)
	insertar(raiz, 1)
	insertar(raiz, 30)
	insertar(raiz, 10)
	insertar(raiz, 2)

	fmt.Printf("\nImprimiendo preorden\n")
	preorden(raiz)
	fmt.Printf("\nImprimiendo inorden\n")
	inorden(raiz)
	fmt.Printf("\nImprimiendo postorden\n")
	postorden(raiz)

	fmt.Printf("\nBuscando el 30\n")
	buscado := buscar(raiz, 30)
	if buscado == nil {
		fmt.Printf("No encontrado!!!!!")
	} else {
		fmt.Printf("Encontrado")
	}
}

