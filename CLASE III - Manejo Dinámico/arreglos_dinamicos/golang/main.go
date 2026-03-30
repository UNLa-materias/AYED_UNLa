package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("-------- Arreglos dinamicos -------\n")

	matriz := crearMatriz(4, 8)
	cargarMatriz(matriz, 4, 8)
	mostrarMatriz(matriz, 4, 8)

	matriz = crearMatriz(11, 3)
	cargarMatriz(matriz, 11, 3)
	mostrarMatriz(matriz, 11, 3)
}

func crearVector(t int) []int {
	return make([]int, t)
}

func cargarVector(v []int, t int) {
	for i := 0; i < t; i++ {
		v[i] = rand.Intn(10)
	}
}

func mostrarVector(v []int, t int) {
	fmt.Printf("\n")
	for i := 0; i < t; i++ {
		fmt.Printf(" %d ", v[i])
	}
}

func crearMatriz(r, c int) [][]int {
	m := make([][]int, r)
	for i := 0; i < r; i++ {
		m[i] = make([]int, c)
	}
	return m
}

func cargarMatriz(m [][]int, r, c int) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m[i][j] = rand.Intn(10)
		}
	}
}

func mostrarMatriz(m [][]int, r, c int) {
	fmt.Printf("\n\n")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf(" %d ", m[i][j])
		}
		fmt.Printf("\n")
	}
}

