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

	// Quiero cambiar algo del carton
	jugador.GetCartonPtr().SetVecPos(2, 99)
	mostrarPersona(jugador)

	vector := [8]int{2, 63, 96, 6, 3, 2, 12, 1}
	jugador.GetCartonPtr().SetVec(vector)
	mostrarPersona(jugador)
}

