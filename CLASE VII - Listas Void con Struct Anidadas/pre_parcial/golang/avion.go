package main

import "fmt"

type Avion struct {
	tipo      string
	matricula int
	capacidad int

	listaPasajeros *Lista[*Pasajero]
}

func crearAvion(tipo string, matricula int, capacidad int) *Avion {
	return &Avion{
		tipo:          tipo,
		matricula:     matricula,
		capacidad:     capacidad,
		listaPasajeros: crearLista[*Pasajero](),
	}
}

func (a *Avion) GetTipo() string                 { return a.tipo }
func (a *Avion) GetMatricula() int               { return a.matricula }
func (a *Avion) GetCapacidad() int               { return a.capacidad }
func (a *Avion) GetListaPasajeros() *Lista[*Pasajero] { return a.listaPasajeros }
func (a *Avion) SetTipo(tipo string)             { a.tipo = tipo }
func (a *Avion) SetMatricula(m int)              { a.matricula = m }
func (a *Avion) SetCapacidad(c int)              { a.capacidad = c }

func agregarPasajero(a *Avion, p *Pasajero) bool {
	if a.capacidad == 0 {
		fmt.Printf("\nNo hay mas asientos disponibles para agregar al pasajero %s\n", p.GetApellido())
		return false
	}
	insertarUltimo(a.listaPasajeros, p)
	a.capacidad--
	return true
}

func eliminarPasajerosVentanilla(a *Avion) {
	actual := getPrimero(a.listaPasajeros)
	anterior := (*Nodo[*Pasajero])(nil)
	pos := 0

	for actual != nil {
		p := actual.dato
		if p.GetVentanilla() == 'S' {
			if anterior == nil {
				_, _ = eliminarPrimero(a.listaPasajeros)
				actual = getPrimero(a.listaPasajeros)
			} else {
				siguiente := actual.siguiente
				_, _ = eliminarPosicion(a.listaPasajeros, pos)
				actual = siguiente
				a.capacidad++
				continue
			}
			a.capacidad++
		} else {
			anterior = actual
			actual = actual.siguiente
			pos++
		}
	}
}

func mostrarAvion(a *Avion) {
	fmt.Printf("\n---------AVION---------\n")
	fmt.Printf("\nTIPO: %s -- MATRICULA: %d -- CAPACIDAD: %d\n", a.tipo, a.matricula, a.capacidad)
	mostrarListaGenerica(a.listaPasajeros, func(p *Pasajero) { mostrarPsajeroDatoPtr(p) })
}

func mostrarAvionDatoPtr(a *Avion) {
	mostrarAvion(a)
}

