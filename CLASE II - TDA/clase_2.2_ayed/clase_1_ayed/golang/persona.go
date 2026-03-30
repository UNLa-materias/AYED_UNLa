package main

import "fmt"

type Persona struct {
	dni    int
	nombre string
	carton *Carton
}

func (p *Persona) GetDNI() int {
	return p.dni
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) GetCartonPtr() *Carton {
	return p.carton
}

func (p *Persona) SetDNI(nuevoDNI int) {
	p.dni = nuevoDNI
}

func (p *Persona) SetNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

func (p *Persona) SetCartonPtr(nuevoCarton *Carton) {
	p.carton = nuevoCarton
}

func crearPersona(dni int, nombre string) *Persona {
	return &Persona{
		dni:    dni,
		nombre: nombre,
		carton: crearCartonAutomatico("Vacio"),
	}
}

func crearPersonaYCarton(dni int, nombre string, dir string) *Persona {
	p := crearPersona(dni, nombre)
	p.carton.SetDireccion(dir)
	return p
}

func mostrarPersona(p *Persona) {
	fmt.Printf("\n---------- PERSONA -----------\n")
	fmt.Printf("DNI: %d\n", p.dni)
	fmt.Printf("NOMBRE: %s\n", p.nombre)
	mostrarCarton(p.carton)
}

