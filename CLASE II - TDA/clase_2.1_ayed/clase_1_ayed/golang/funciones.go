package main

import "math/rand"

func aleatorioEntre(minimo, maximo int) int {
	return minimo + rand.Intn(maximo-minimo+1)
}

func busquedaVectorInt(vec []int, busco int) int {
	for i, v := range vec {
		if busco == v {
			return i
		}
	}
	return -1
}

func cargarVectorIntConValor(vec []int, valor int) {
	for i := range vec {
		vec[i] = valor
	}
}

func cargarVectorIntAleatorio(vec []int, minimo, maximo int) {
	for i := range vec {
		vec[i] = aleatorioEntre(minimo, maximo)
	}
}

func cargarVectorIntAleatorioSinRepetir(vec []int, minimo, maximo int) {
	for i := range vec {
		var numero int
		for {
			numero = aleatorioEntre(minimo, maximo)
			if busquedaVectorInt(vec, numero) == -1 {
				break
			}
		}
		vec[i] = numero
	}
}

