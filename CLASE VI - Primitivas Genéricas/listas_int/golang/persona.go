package main

import "fmt"

type Persona struct {
	dni    int
	nombre string
}

func crearPersona(dni int, nombre string) *Persona {
	return &Persona{dni: dni, nombre: nombre}
}

func mostrarPersona(p *Persona) {
	fmt.Printf("\n---------- PERSONA -----------\n")
	fmt.Printf("DNI: %d\n", p.dni)
	fmt.Printf("NOMBRE: %s\n", p.nombre)
}

func mostrarPersonaDatoPtr(p *Persona) {
	mostrarPersona(p)
}

func compararPersonasDatoPtr(p1, p2 *Persona) bool {
	return p1.dni > p2.dni
}

