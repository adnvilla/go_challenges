package controllers

import (
	"fmt"
	"net/http"

	usecases "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	errorshandle "github.com/adnvilla/go_challenges/investment_api/src/pkg/errors_handle"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	usecase use_case.UseCase[usecases.GetStatisticsInput, usecases.GetStatisticsOutput]
}

func NewStatisticsHandlerHandler(usecase use_case.UseCase[usecases.GetStatisticsInput, usecases.GetStatisticsOutput]) StatisticsHandler {
	return StatisticsHandler{
		usecase: usecase,
	}
}

func (handler *StatisticsHandler) GetStatistics(c *gin.Context) {
	input := usecases.GetStatisticsInput{}
	result, err := handler.usecase.Handle(c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result.StatisticsReponse)
}
