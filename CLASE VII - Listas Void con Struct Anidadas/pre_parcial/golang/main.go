package main

import "fmt"

func main() {
	fmt.Printf("\n-----------EJEMPLO PRE PARCIAL --------------\n\n")

	p1 := crearPasajero("Apellido1", 1, 'S')
	p2 := crearPasajero("Apellido2", 2, 'S')
	p3 := crearPasajero("Apellido3", 3, 'N')
	p4 := crearPasajero("Apellido4", 4, 'S')
	p5 := crearPasajero("Apellido5", 5, 'N')
	p6 := crearPasajero("Apellido6", 6, 'S')

	avion1 := crearAvion("Boeing 737", 1111, 10)
	avion2 := crearAvion("Airbus A320", 2222, 5)
	avion3 := crearAvion("Embraer E190", 3333, 5)

	aeropuerto := crearAeropuerto("Ministro Pistarini", "Riccheri 33", "11-5480-61111")

	insertarAvion(aeropuerto, avion1)
	insertarAvion(aeropuerto, avion2)
	insertarAvion(aeropuerto, avion3)

	fmt.Printf("\n\n------ Aeropuerto SIN  pasajeros ----\n\n")
	mostrarAeropuerto(aeropuerto)

	agregarPasajero(avion1, p1)
	agregarPasajero(avion1, p2)
	agregarPasajero(avion1, p3)
	agregarPasajero(avion2, p4)
	agregarPasajero(avion2, p5)
	agregarPasajero(avion3, p6)

	fmt.Printf("\n\n------ Aeropuerto ya con pasajeros ----\n\n")
	mostrarAeropuerto(aeropuerto)
}

