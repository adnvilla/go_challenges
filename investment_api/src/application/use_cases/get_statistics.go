package use_cases

import (
	"context"

	"github.com/adnvilla/go_challenges/investment_api/src/domain/repository"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/dto"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/use_case"
)

type GetStatisticsUseCase struct {
	repository repository.CreditAssignmentRepository
}

type GetStatisticsInput struct {
}

type GetStatisticsOutput struct {
	StatisticsReponse dto.Statistic
}

func NewGetStatisticsUseCase(r repository.CreditAssignmentRepository) use_case.UseCase[GetStatisticsInput, GetStatisticsOutput] {
	u := new(GetStatisticsUseCase)
	u.repository = r
	return u
}

func (u *GetStatisticsUseCase) Handle(ctx context.Context, input GetStatisticsInput) (GetStatisticsOutput, error) {
	statistics, err := u.repository.GetStatistics()
	if err != nil {
		return GetStatisticsOutput{}, err
	}

	// This wil be moved to mapper dtos
	output := GetStatisticsOutput{}
	totalAssignmentSuccess := 0
	totalAssignmentUnSuccess := 0
	TotalAmountgAssignmentSuccess := 0
	TotalAmountAssignmentUnSuccess := 0
	for i := 0; i < len(statistics); i++ {
		currentSt := statistics[i]
		if currentSt.Success {
			totalAssignmentSuccess++
			TotalAmountgAssignmentSuccess += currentSt.Investment
		} else {
			totalAssignmentUnSuccess++
			TotalAmountAssignmentUnSuccess += currentSt.Investment
		}
	}

	if len(statistics) > 0 {
		output.StatisticsReponse = dto.Statistic{
			TotalAssignments:         len(statistics),
			TotalAssignmentSuccess:   totalAssignmentSuccess,
			TotalAssignmentUnSuccess: totalAssignmentUnSuccess,
		}

		if totalAssignmentSuccess > 0 {
			output.StatisticsReponse.AverageAssignmentSuccess = float32(TotalAmountgAssignmentSuccess) / float32(totalAssignmentSuccess)
		}
		if totalAssignmentUnSuccess > 0 {
			output.StatisticsReponse.AverageAssignmentUnSuccess = float32(TotalAmountAssignmentUnSuccess) / float32(totalAssignmentUnSuccess)
		}
	}
	return output, nil
}
