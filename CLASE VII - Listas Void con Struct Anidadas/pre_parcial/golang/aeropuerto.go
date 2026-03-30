package main

import "fmt"

type Aeropuerto struct {
	nombre    string
	direccion string
	telefono  string

	listaAviones *Lista[*Avion]
}

func crearAeropuerto(nombre, direccion, telefono string) *Aeropuerto {
	return &Aeropuerto{
		nombre:      nombre,
		direccion:   direccion,
		telefono:    telefono,
		listaAviones: crearLista[*Avion](),
	}
}

func (a *Aeropuerto) GetNombre() string          { return a.nombre }
func (a *Aeropuerto) GetDireccion() string       { return a.direccion }
func (a *Aeropuerto) GetTelefono() string        { return a.telefono }
func (a *Aeropuerto) SetNombre(nombre string)    { a.nombre = nombre }
func (a *Aeropuerto) SetDireccion(dir string)    { a.direccion = dir }
func (a *Aeropuerto) SetTelefono(tel string)     { a.telefono = tel }
func (a *Aeropuerto) GetListaAviones() *Lista[*Avion] { return a.listaAviones }

func insertarAvion(a *Aeropuerto, av *Avion) {
	insertarUltimo(a.listaAviones, av)
}

func mostrarAeropuerto(a *Aeropuerto) {
	fmt.Printf("\n---------AEROPUERTO---------\n")
	fmt.Printf("\nNOMBRE: %s -- DIRECCION: %s -- TELEFONO: %s\n", a.nombre, a.direccion, a.telefono)
	mostrarListaGenerica(a.listaAviones, func(av *Avion) { mostrarAvionDatoPtr(av) })
}

