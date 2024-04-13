package main

import (
	"fmt"
	"sort"
)

// Función para encontrar todas las combinaciones posibles de sumas de los números primos dados, permitiendo duplicación
func findCombinations(primes []int, target int) int {
	count := 0
	//[]int{3, 5, 7} Se establece para que al menos este presente cada numero primo en la combinacion que se encuentre
	// Si desea que se permita combinaciones que satisfagan la suma con un solo numero primo, establecer []int{}
	combinations(primes, target, []int{3, 5, 7}, &count)
	return count
}

// Función recursiva para generar todas las combinaciones posibles con duplicación
func combinations(primes []int, target int, currentComb []int, count *int) {
	sum := 0
	for _, num := range currentComb {
		sum += num
	}
	if sum == target {
		fmt.Println(currentComb) // printSortedComb para revisar las combinaciones
		*count++
		return
	}
	if sum > target {
		return
	}
	for i := 0; i < len(primes); i++ {
		combinations(primes, target, append(currentComb, primes[i]), count)
		if *count > 0 {
			return // Si ya se encontró una combinación válida, terminar la verificación de combinaciones
		}
	}
}

// Función para imprimir la combinación ordenada y verificar si ya se ha impreso
func printSortedComb(comb []int) {
	sort.Ints(comb)
	staticCombinations := make(map[string]bool)
	sortedComb := fmt.Sprintf("%v", comb)
	if !staticCombinations[sortedComb] {
		fmt.Println(comb)
		staticCombinations[sortedComb] = true
	}
}

func main() {
	// Ya que se garantiza que son multiplos de 100, entonces podemos inferir que los valores 300, 500, y 700 son los numeros primos de 3, 5 y 7
	primes := []int{3, 5, 7}
	targetNumber := 30
	countCombinations := findCombinations(primes, targetNumber)
	fmt.Println("Número total de combinaciones:", countCombinations)

	// Para obtener todas las combinaciones, quitar la restriccion de cuando encuentre una combinacon que satisfaga las condiciones regrese esta combinacion
}

var _ CreditAssigner = (*CreditAssignment)(nil)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignment struct {
}

func (*CreditAssignment) Assign(investment int32) (x300 int32, x500 int32, x700 int32, err error) {
	return
}
