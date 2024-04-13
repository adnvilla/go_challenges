package main

import (
	"slices"
	"testing"

	"github.com/adnvilla/go_challenges/investment_api/algorithm"
)

func TestInvestment(t *testing.T) {

	investment := []struct {
		tarjet        int
		findExpected  bool
		expectedCombi []int
	}{
		{30, true, []int{3, 3, 5, 5, 7, 7}},
		{67, true, []int{3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 7, 7, 7}},
		{4, false, []int{}},
	}

	primes := []int{3, 5, 7}
	initComb := []int{3, 5, 7}

	for _, inv := range investment {
		_, combis := algorithm.FindCombinations(primes, inv.tarjet, initComb, false)

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
