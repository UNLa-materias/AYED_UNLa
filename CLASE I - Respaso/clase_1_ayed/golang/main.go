package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("------ Clase 1 - AYED - REPASO -------\n")

	jugador := crearPersona(123, "Nico")
	mostrarPersona(jugador)

	carton := crearCarton1("Av. Peron 1130")
	mostrarCarton(carton)
}

