package repository

import domain "github.com/adnvilla/go_challenges/investment_api/src/domain/entity"

type CreditAssignmentRepository interface {
	SaveStatistics(domain.Statistic) error
	GetStatistics() ([]domain.Statistic, error)
}
