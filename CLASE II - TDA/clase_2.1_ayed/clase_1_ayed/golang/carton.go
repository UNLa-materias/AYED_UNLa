package main

import "fmt"

type Carton struct {
	vec       [8]int
	direccion string
}

func crearCartonAutomatico(direccion string) *Carton {
	c := &Carton{direccion: direccion}

	tmp := c.vec[:]
	cargarVectorIntConValor(tmp, -1)
	// En el C: se usa 0..10 para facilitar ver repetidos.
	cargarVectorIntAleatorioSinRepetir(tmp, 0, 10)

	return c
}

func mostrarCarton(c *Carton) {
	fmt.Printf("\n\t---------- CARTON ---------\n")
	fmt.Printf("\tDIRECCION: %s\n\t---VECTOR---\n\t", c.direccion)
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d ", c.vec[i])
	}
}

