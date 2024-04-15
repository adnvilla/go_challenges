package infrastructure

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/adnvilla/go_challenges/investment_api/src/domain/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetStatistics(t *testing.T) {
	// Mock DB
	db, dbMock := GetMockDB()
	r, _ := db.DB()
	defer r.Close()
	//Fixture
	sqlRegex := `SELECT * FROM "statistics"`
	success := true
	investment := 3000

	expected := []domain.Statistic{
		{
			Success:    true,
			Investment: 3000,
		},
	}

	dbMock.ExpectQuery(regexp.QuoteMeta(sqlRegex)).WillReturnRows(
		sqlmock.NewRows([]string{"success", "investment"}).AddRow(
			success,
			investment,
		))
	repository := NewCreditAssignmentRepository(db)

	// Act
	response, err := repository.GetStatistics()

	// Assert
	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetStatisticsFail(t *testing.T) {
	// Mock DB
	db, dbMock := GetMockDB()
	r, _ := db.DB()
	defer r.Close()
	//Fixture
	sqlRegex := `SELECT * FROM "statistics"`

	expected := []domain.Statistic{}

	dbMock.ExpectQuery(regexp.QuoteMeta(sqlRegex)).WillReturnError(fmt.Errorf("some error"))
	repository := NewCreditAssignmentRepository(db)

	// Act
	response, err := repository.GetStatistics()

	// Assert
	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Error(t, err)
	assert.Equal(t, expected, response)
}

func TestSaveStatistics(t *testing.T) {
	// Mock DB
	db, dbMock := GetMockDB()
	r, _ := db.DB()
	defer r.Close()
	//Fixture
	sqlRegex := `INSERT INTO "statistics"`
	request := domain.Statistic{
		Success:    true,
		Investment: 3000,
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).WithArgs(request.Success, request.Investment).WillReturnResult(sqlmock.NewResult(1, 1))
	dbMock.ExpectCommit()
	repository := NewCreditAssignmentRepository(db)

	// Act
	err := repository.SaveStatistics(request)

	// Assert
	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NoError(t, err)
}

func TestSaveStatisticsFail(t *testing.T) {
	// Mock DB
	db, dbMock := GetMockDB()
	r, _ := db.DB()
	defer r.Close()
	//Fixture
	sqlRegex := `INSERT INTO "statistics"`
	request := domain.Statistic{
		Success:    true,
		Investment: 3000,
	}
	dbMock.ExpectBegin()
	dbMock.ExpectExec(sqlRegex).
		WithArgs(request.Success, request.Investment).
		WillReturnError(fmt.Errorf("some error"))
	dbMock.ExpectRollback()
	repository := NewCreditAssignmentRepository(db)

	// Act
	err := repository.SaveStatistics(request)

	// Assert
	if err := dbMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Error(t, err)
}

func GetMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	// Mock DB
	db, dbMock, _ := sqlmock.New()

	// Use postgresql to avoid generate sqlite files
	gdb, _ := gorm.Open(postgres.New(
		postgres.Config{
			Conn:       db,
			DriverName: "postgres",
		},
	), &gorm.Config{})

	return gdb, dbMock
}
