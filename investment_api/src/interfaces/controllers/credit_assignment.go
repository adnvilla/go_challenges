package controllers

import (
	"errors"
	"fmt"
	"net/http"

	usecases "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/dto"
	errorshandle "github.com/adnvilla/go_challenges/investment_api/src/pkg/errors_handle"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type CreditAssigmentHandler struct {
	usecase use_case.UseCase[usecases.CreateCreditAssignmentInput, usecases.CreateCreditAssignmentOutput]
}

func NewCreateCreditAssignmentHandler(usecase use_case.UseCase[usecases.CreateCreditAssignmentInput, usecases.CreateCreditAssignmentOutput]) CreditAssigmentHandler {
	return CreditAssigmentHandler{
		usecase: usecase,
	}
}

func (handler *CreditAssigmentHandler) CreateCreditAssignment(c *gin.Context) {
	var body dto.CreditAssignmentRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(errors.New("please check the request")),
		})
		return
	}

	input := usecases.CreateCreditAssignmentInput{
		Investment: body.Investment,
	}

	result, err := handler.usecase.Handle(c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result.CreditAssignmentResponse)
}
