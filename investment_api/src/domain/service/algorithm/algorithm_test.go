package algorithm

import (
	"fmt"
	"slices"
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

func TestInvestment(t *testing.T) {

	investment := []struct {
		tarjet        int
		findExpected  bool
		expectedCombi []int
		findOnce      bool
	}{
		{30, true, []int{3, 3, 5, 5, 7, 7}, false},
		{30, true, []int{3, 3, 3, 3, 3, 3, 5, 7}, true},
		{67, true, []int{3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 7, 7, 7}, false},
		{4, false, []int{}, false},
	}

	primes := []int{3, 5, 7}
	initComb := []int{3, 5, 7}

	for _, inv := range investment {
		_, combis := FindCombinations(primes, inv.tarjet, initComb, inv.findOnce)

		find := false
		for i := 0; i < len(combis); i++ {
			currentComb := combis[i]
			slices.Sort(currentComb)
			if testEq(inv.expectedCombi, currentComb) {
				find = true
				break
			}
		}

		if !find && inv.findExpected {
			t.Errorf("The case tarjet %v dont find the expected combi %v", inv.tarjet, inv.expectedCombi)
		}

		if find && !inv.findExpected {
			t.Errorf("The case tarjet %v find the expected combi %v", inv.tarjet, inv.expectedCombi)
		}
	}
}

func TestInvestmentWithGoroutines(t *testing.T) {

	investment := []struct {
		tarjet        int
		findExpected  bool
		expectedCombi []int
		findOnce      bool
	}{
		{30, true, []int{3, 3, 5, 5, 7, 7}, false},
		{30, true, []int{3, 3, 3, 3, 3, 3, 5, 7}, true},
		{67, true, []int{3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 7, 7, 7}, false},
		{4, false, []int{}, false},
	}

	primes := []int{3, 5, 7}
	initComb := []int{3, 5, 7}

	for _, inv := range investment {
		_, combis := FindCombinationsWithGoroutines(primes, inv.tarjet, initComb, inv.findOnce)

		find := false
		for i := 0; i < len(combis); i++ {
			currentComb := combis[i]
			slices.Sort(currentComb)

			if testEq(inv.expectedCombi, currentComb) {
				find = true
				break
			}
		}

		if !find && inv.findExpected {
			t.Errorf("The case tarjet %v dont find the expected combi %v", inv.tarjet, inv.expectedCombi)
		}

		if find && !inv.findExpected {
			t.Errorf("The case tarjet %v find the expected combi %v", inv.tarjet, inv.expectedCombi)
		}
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
