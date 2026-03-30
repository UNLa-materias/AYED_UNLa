package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("------ Clase 1 - AYED - REPASO -------\n")

	jugador := crearPersonaYCarton(123, "Nico", "Av. Peron 1130")
	mostrarPersona(jugador)
}

