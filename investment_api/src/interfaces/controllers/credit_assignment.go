package controllers

import (
	"errors"
	"fmt"
	"net/http"

	usecases "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/service"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/dto"
	errorshandle "github.com/adnvilla/go_challenges/investment_api/src/pkg/errors_handle"
	"github.com/gin-gonic/gin"
)

func CreateCreditAssignment(c *gin.Context) {
	// This can be moved to dispatcher
	usecase := usecases.NewCreateCreditAssignmentUseCase(service.NewCreditAssignmentService())

	// get dl body
	var dlBody dto.CreditAssignmentRequest
	err := c.ShouldBindJSON(&dlBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(errors.New("please check the request")),
		})
		return
	}

	input := usecases.CreateCreditAssignmentInput{
		Investment: dlBody.Investment,
	}

	result, err := usecase.Handle(c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
