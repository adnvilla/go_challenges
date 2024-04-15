package infrastructure

import (
	"errors"

	domain "github.com/adnvilla/go_challenges/investment_api/src/domain/entity"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/repository"
	"github.com/adnvilla/go_challenges/investment_api/src/infrastructure/models"
	"gorm.io/gorm"
)

type creditAssignmentRepository struct {
	db *gorm.DB
}

func NewCreditAssignmentRepository(db *gorm.DB) repository.CreditAssignmentRepository {
	return &creditAssignmentRepository{
		db: db,
	}
}

func (r *creditAssignmentRepository) SaveStatistics(s domain.Statistic) error {
	statistic := models.Statistic{
		Success:    s.Success,
		Investment: s.Investment}
	result := r.db.Create(&statistic)
	if result.Error != nil {
		return errors.New("have a issue with consult DB")
	}

	return nil
}
func (r *creditAssignmentRepository) GetStatistics() ([]domain.Statistic, error) {
	st := []domain.Statistic{}
	statistics := []models.Statistic{}
	result := r.db.Find(&statistics)
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
