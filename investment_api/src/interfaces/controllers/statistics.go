package controllers

import (
	"fmt"
	"net/http"

	usecases "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	"github.com/adnvilla/go_challenges/investment_api/src/infrastructure"
	errorshandle "github.com/adnvilla/go_challenges/investment_api/src/pkg/errors_handle"
	"github.com/gin-gonic/gin"
)

func GetStatistics(c *gin.Context) {
	// This can be moved to dispatcher
	usecase := usecases.NewGetStatisticsUseCase(infrastructure.NewCreditAssignmentRepository())

	input := usecases.GetStatisticsInput{}
	result, err := usecase.Handle(c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result.StatisticsReponse)
}
