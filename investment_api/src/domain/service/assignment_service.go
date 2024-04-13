package service

import (
	"errors"

	"github.com/adnvilla/go_challenges/investment_api/src/domain/service/algorithm"
)

var _ CreditAssigner = (*CreditAssignment)(nil)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssignment struct {
}

func NewCreditAssignmentService() CreditAssigner {
	return new(CreditAssignment)
}

func (*CreditAssignment) Assign(investment int32) (x300 int32, x500 int32, x700 int32, err error) {
	count, combi := algorithm.FindCombinations([]int{3, 5, 7}, int(investment)/100, []int{3, 5, 7}, true)
	if count > 0 {
		r := combi[0]
		for i := 0; i < len(r); i++ {
			switch r[i] {
			case 3:
				x300++
			case 5:
				x500++
			case 7:
				x700++
			default:
				return 0, 0, 0, errors.New("error: can not assignment the investment")
			}
		}
		return
	}
	return 0, 0, 0, errors.New("error: can not assignment the investment")
}
