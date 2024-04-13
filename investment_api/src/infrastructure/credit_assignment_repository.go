package infrastructure

import (
	"errors"

	domain "github.com/adnvilla/go_challenges/investment_api/src/domain/entity"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/repository"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/gorm"
)

type Statistic struct {
	Success    bool `gorm:"column:success"`
	Investment int  `gorm:"column:investment"`
}

type creditAssignmentRepository struct {
	tableName string
}

func NewCreditAssignmentRepository() repository.CreditAssignmentRepository {
	return &creditAssignmentRepository{}
}

func (r *creditAssignmentRepository) SaveStatistics(s domain.Statistic) error {
	db := gorm.GetConnection()

	statistic := Statistic{
		Success:    s.Success,
		Investment: s.Investment}
	result := db.Create(&statistic)
	// SELECT * FROM statistics;
	if result.Error != nil {
		return errors.New("have a issue with consult DB")
	}

	return nil
}
func (r *creditAssignmentRepository) GetStatistics() ([]domain.Statistic, error) {
	db := gorm.GetConnection()

	st := []domain.Statistic{}
	statistics := []Statistic{}
	result := db.Find(&statistics)
	// SELECT * FROM statistics;
	if result.Error != nil {
		return st, errors.New("have a issue with consult DB")
	}

	for i := 0; i < len(statistics); i++ {
		st = append(st, domain.Statistic{
			Success:    statistics[i].Success,
			Investment: statistics[i].Investment,
		})
	}

	return st, nil
}
