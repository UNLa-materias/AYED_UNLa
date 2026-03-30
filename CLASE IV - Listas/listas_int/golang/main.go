package main

func main() {
	lista := crearLista()
	mostrarLista(lista)

	insertarPrimero(lista, 4)
	insertarPrimero(lista, 1)
	insertarPrimero(lista, -26)
	insertarPrimero(lista, 96)

	mostrarLista(lista)
}

