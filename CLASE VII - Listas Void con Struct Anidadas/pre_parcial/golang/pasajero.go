package main

import "fmt"

type Pasajero struct {
	apellido   string
	dni        int
	ventanilla rune
}

func crearPasajero(apellido string, dni int, ventanilla rune) *Pasajero {
	return &Pasajero{apellido: apellido, dni: dni, ventanilla: ventanilla}
}

func (p *Pasajero) GetApellido() string  { return p.apellido }
func (p *Pasajero) GetDni() int          { return p.dni }
func (p *Pasajero) GetVentanilla() rune  { return p.ventanilla }
func (p *Pasajero) SetApellido(a string) { p.apellido = a }
func (p *Pasajero) SetDni(d int)         { p.dni = d }
func (p *Pasajero) SetVentanilla(v rune) { p.ventanilla = v }

func mostrarPasajero(p *Pasajero) {
	fmt.Printf("\n---------PASAJERO---------\n")
	fmt.Printf("\nAPELLIDO: %s\n", p.apellido)
	fmt.Printf("\nDNI: %d\n", p.dni)
	fmt.Printf("\nVENTANILLA: %c\n", p.ventanilla)
}

func mostrarPsajeroDatoPtr(p *Pasajero) {
	mostrarPasajero(p)
}

