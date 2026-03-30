package main

import "fmt"

type Persona struct {
	dni    int
	nombre string

	cartonDeLaPersona *Carton
}

func crearPersona(dni int, nombre string) *Persona {
	return &Persona{
		dni:               dni,
		nombre:            nombre,
		cartonDeLaPersona: crearCartonAutomatico("Vacio"),
	}
}

func crearPersonaYCarton(dni int, nombre string, dir string) *Persona {
	p := crearPersona(dni, nombre)
	p.cartonDeLaPersona.direccion = dir
	return p
}

func mostrarPersona(p *Persona) {
	fmt.Printf("\n---------- PERSONA -----------\n")
	fmt.Printf("DNI: %d\n", p.dni)
	fmt.Printf("NOMBRE: %s\n", p.nombre)
	mostrarCarton(p.cartonDeLaPersona)
}

