package main

import (
	"fmt"
	"math/rand"
)

type Carton struct {
	vec       [8]int
	direccion string
}

func crearCarton1(direccion string) *Carton {
	c := &Carton{direccion: direccion}
	for i := 0; i < 8; i++ {
		c.vec[i] = rand.Intn(100) // 0..99
		// En el C original: falta verificar repetidos; se replica el comportamiento.
	}
	return c
}

func mostrarCarton(c *Carton) {
	fmt.Printf("\n\t------ CARTON ----\n")
	fmt.Printf("\tDIRECCION: %s\n\t---VECTOR---\n\t", c.direccion)
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d ", c.vec[i])
	}
}

