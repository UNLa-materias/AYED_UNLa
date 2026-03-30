package main

import "fmt"

func main() {
	var unArbol *ArbolAVL

	insertar(&unArbol, 1)
	insertar(&unArbol, 55)
	insertar(&unArbol, 24)
	insertar(&unArbol, 12)
	insertar(&unArbol, 13)

	mostrar(unArbol)
	fmt.Printf("\n\n")
}

