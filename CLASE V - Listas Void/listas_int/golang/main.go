package main

func main() {
	lista := crearLista()
	mostrarLista(lista)

	insertarPrimero(lista, 4)
	insertarPrimero(lista, 11)
	insertarUltimo(lista, 109)
	insertarPrimero(lista, 21)
	insertarPrimero(lista, 23)
	insertarUltimo(lista, 300)
	insertarPrimero(lista, 34)

	mostrarLista(lista)

	ordenarLista(lista)
	mostrarLista(lista)
}

