package main

import (
	"fmt"

	"github.com/adnvilla/go_challenges/investment_api/algorithm"
)

func main() {
	// Ya que se garantiza que son multiplos de 100, entonces podemos inferir que los valores 300, 500, y 700 son los numeros primos de 3, 5 y 7
	primes := []int{3, 5, 7}
	targetNumber := 670
	countCombinations, combis := algorithm.FindCombinations(primes, targetNumber, []int{3, 5, 7}, false)
	fmt.Println("Número total de combinaciones:", countCombinations)
	fmt.Println(combis)

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
