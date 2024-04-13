package use_cases

import (
	"context"

	domain "github.com/adnvilla/go_challenges/investment_api/src/domain/entity"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/repository"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/service"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/dto"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/use_case"
)

type CreateCreditAssignmentUseCase struct {
	service    service.CreditAssigner
	repository repository.CreditAssignmentRepository
}

type CreateCreditAssignmentInput struct {
	Investment int
}

type CreateCreditAssignmentOutput struct {
	CreditAssignmentResponse dto.CreditAssignmentResponse
}

func NewCreateCreditAssignmentUseCase(s service.CreditAssigner, r repository.CreditAssignmentRepository) use_case.UseCase[CreateCreditAssignmentInput, CreateCreditAssignmentOutput] {
	u := new(CreateCreditAssignmentUseCase)
	u.service = s
	u.repository = r
	return u
}

func (u *CreateCreditAssignmentUseCase) Handle(ctx context.Context, input CreateCreditAssignmentInput) (CreateCreditAssignmentOutput, error) {
	x300, x500, x700, err := u.service.Assign(int32(input.Investment))

	if err != nil {
		u.repository.SaveStatistics(domain.Statistic{
			Success:    false,
			Investment: input.Investment,
		})
		return CreateCreditAssignmentOutput{}, err
	}

	assign := domain.CreditAssignment{
		Credit300: int(x300),
		Credit500: int(x500),
		Credit700: int(x700),
	}

	u.repository.SaveStatistics(domain.Statistic{
		Success:    true,
		Investment: input.Investment,
	})

	// This wil be moved to mapper dtos
	output := CreateCreditAssignmentOutput{
		CreditAssignmentResponse: dto.CreditAssignmentResponse{
			Credit300: assign.Credit300,
			Credit500: assign.Credit500,
			Credit700: assign.Credit700,
		},
	}

	return output, nil
}
