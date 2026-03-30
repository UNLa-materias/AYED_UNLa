package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cargarDosArrayIdenticos(array1, array2 []int) {
	for i := 0; i < len(array1); i++ {
		num := rand.Intn(1000)
		array1[i] = num
		array2[i] = num
	}
}

func encontrarParFuerzaBruta(arr []int, t int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == t {
				fmt.Printf("\n----------- FORx2 -----------")
				fmt.Printf("\n\nPar encontrado: (%d, %d)\n", arr[i], arr[j])
				return true
			}
		}
	}
	fmt.Printf("No se encontro ningun par que sume %d\n", t)
	return false
}

func encontrarParHashMap(arr []int, t int) bool {
	vistos := make(map[int]struct{}, len(arr))
	for i := 0; i < len(arr); i++ {
		complemento := t - arr[i]
		if _, ok := vistos[complemento]; ok {
			fmt.Printf("\n----------- HASHMAP -----------")
			fmt.Printf("\n\nPar encontrado: (%d, %d)\n", complemento, arr[i])
			return true
		}
		vistos[arr[i]] = struct{}{}
	}
	fmt.Printf("No se encontró ningún par que sume %d\n", t)
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	loQueBusco := 9

	n := 20000
	vector1 := make([]int, n)
	vector2 := make([]int, n)

	cargarDosArrayIdenticos(vector1, vector2)

	start := time.Now()
	encontrarParFuerzaBruta(vector1, loQueBusco)
	fmt.Printf("\nEl algoritmo doble FOR tardo %f milisegundos en ejecutarse.\n", float64(time.Since(start).Microseconds())/1000.0)

	start = time.Now()
	encontrarParHashMap(vector2, loQueBusco)
	fmt.Printf("\nEl algoritmo  HashMap tardo %f milisegundos en ejecutarse.\n", float64(time.Since(start).Microseconds())/1000.0)
}

