package algorithm

import (
	"fmt"
	"testing"
)

func BenchmarkFindCombinations30(b *testing.B) {
	primes := []int{3, 5, 7}
	targetNumber := 30
	c, _ := FindCombinations(primes, targetNumber, []int{3, 5, 7}, false)
	fmt.Println(c)
}

func BenchmarkFindCombinations67(b *testing.B) {
	primes := []int{3, 5, 7}
	targetNumber := 67
	c, _ := FindCombinations(primes, targetNumber, []int{3, 5, 7}, false)
	fmt.Println(c)
}

func BenchmarkFindCombinationsWithGoroutinesr30(b *testing.B) {
	primes := []int{3, 5, 7}
	targetNumber := 67
	c, _ := FindCombinations(primes, targetNumber, []int{3, 5, 7}, false)
	fmt.Println(c)
}

func BenchmarkFindCombinationsWithGoroutines67(b *testing.B) {
	primes := []int{3, 5, 7}
	targetNumber := 30
	c, _ := FindCombinationsWithGoroutines(primes, targetNumber, []int{3, 5, 7}, false)
	fmt.Println(c)
}
