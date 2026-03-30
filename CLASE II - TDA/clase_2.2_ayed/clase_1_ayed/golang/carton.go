package main

import "fmt"

type Carton struct {
	vec       [8]int
	direccion string
}

func (c *Carton) GetDireccion() string {
	return c.direccion
}

func (c *Carton) GetVec() [8]int {
	return c.vec
}

func (c *Carton) SetDireccion(nuevaDireccion string) {
	c.direccion = nuevaDireccion
}

func (c *Carton) SetVec(nuevoVec [8]int) {
	c.vec = nuevoVec
}

func (c *Carton) GetVecPos(pos int) int {
	return c.vec[pos]
}

func (c *Carton) SetVecPos(pos int, nuevoElemento int) {
	c.vec[pos] = nuevoElemento
}

func crearCartonAutomatico(direccion string) *Carton {
	c := &Carton{direccion: direccion}
	tmp := c.vec[:]
	cargarVectorIntConValor(tmp, -1)
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

