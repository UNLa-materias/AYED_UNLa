package main

func main() {
	p1 := crearPersona(1, "Nico")
	p2 := crearPersona(33, "Walter")
	p3 := crearPersona(12, "Brenda")

	lista := crearLista[*Persona]()

	insertarPrimero(lista, p1)
	insertarPrimero(lista, p2)
	insertarPrimero(lista, p3)

	mostrarListaGenerica(lista, func(p *Persona) { mostrarPersonaDatoPtr(p) })

	ordenarLista(lista, func(a, b *Persona) bool { return compararPersonasDatoPtr(a, b) })

	mostrarListaGenerica(lista, func(p *Persona) { mostrarPersonaDatoPtr(p) })
}

