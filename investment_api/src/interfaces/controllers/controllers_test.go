package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	usecase "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	usecasemock "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases/mock"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func mockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	if content == nil {
		return
	}

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestMockJson(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)
	mockJsonPost(ctx, nil)
}

func TestCreateCreditAssignment(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.CreditAssignmentRequest{
		Investment: 3000,
	}
	expected := dto.CreditAssignmentResponse{
		Credit300: 6,
		Credit500: 1,
		Credit700: 1,
	}
	expectedStatus := http.StatusOK
	response := dto.CreditAssignmentResponse{}

	mockJsonPost(ctx, body)

	usecaseMock := usecasemock.NewMockCreateCreditAssignmentUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateCreditAssignmentOutput{
		CreditAssignmentResponse: dto.CreditAssignmentResponse{
			Credit300: 6,
			Credit500: 1,
			Credit700: 1,
		},
	}, nil)

	// Act
	handlerStatistics := NewCreateCreditAssignmentHandler(usecaseMock)
	handlerStatistics.CreateCreditAssignment(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateCreditAssignmentFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.CreditAssignmentRequest{
		Investment: 3000,
	}
	expected := dto.CreditAssignmentResponse{
		Credit300: 0,
		Credit500: 0,
		Credit700: 0,
	}
	expectedStatus := http.StatusBadRequest
	response := dto.CreditAssignmentResponse{}

	mockJsonPost(ctx, body)

	usecaseMock := usecasemock.NewMockCreateCreditAssignmentUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateCreditAssignmentOutput{}, fmt.Errorf("some error"))

	// Act
	handlerStatistics := NewCreateCreditAssignmentHandler(usecaseMock)
	handlerStatistics.CreateCreditAssignment(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateCreditAssignmentFailBody(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture

	expected := dto.CreditAssignmentResponse{
		Credit300: 0,
		Credit500: 0,
		Credit700: 0,
	}
	expectedStatus := http.StatusBadRequest
	response := dto.CreditAssignmentResponse{}

	mockJsonPost(ctx, nil)

	usecaseMock := usecasemock.NewMockCreateCreditAssignmentUseCase(t)
	//usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateCreditAssignmentOutput{}, fmt.Errorf("some error"))

	// Act
	handlerStatistics := NewCreateCreditAssignmentHandler(usecaseMock)
	handlerStatistics.CreateCreditAssignment(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetStatistics(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture
	expected := dto.Statistic{
		TotalAssignments:           1,
		TotalAssignmentSuccess:     1,
		TotalAssignmentUnSuccess:   0,
		AverageAssignmentSuccess:   0,
		AverageAssignmentUnSuccess: 0,
	}
	expectedStatus := http.StatusOK
	response := dto.Statistic{}

	mockJsonPost(ctx, nil)

	usecaseMock := usecasemock.NewMockGetStatisticsUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetStatisticsOutput{StatisticsReponse: dto.Statistic{
		TotalAssignments:           1,
		TotalAssignmentSuccess:     1,
		TotalAssignmentUnSuccess:   0,
		AverageAssignmentSuccess:   0,
		AverageAssignmentUnSuccess: 0,
	}}, nil)

	// Act
	handlerStatistics := NewStatisticsHandlerHandler(usecaseMock)

	handlerStatistics.GetStatistics(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetStatisticsFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture
	expected := dto.Statistic{
		TotalAssignments:           0,
		TotalAssignmentSuccess:     0,
		TotalAssignmentUnSuccess:   0,
		AverageAssignmentSuccess:   0,
		AverageAssignmentUnSuccess: 0,
	}
	expectedStatus := http.StatusBadRequest
	response := dto.Statistic{}

	mockJsonPost(ctx, nil)

	usecaseMock := usecasemock.NewMockGetStatisticsUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetStatisticsOutput{}, fmt.Errorf("some error"))

	// Act
	handlerStatistics := NewStatisticsHandlerHandler(usecaseMock)

	handlerStatistics.GetStatistics(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}
