package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	"github.com/adnvilla/go_challenges/investment_api/src/domain/service"
	"github.com/adnvilla/go_challenges/investment_api/src/infrastructure"
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/controllers"
	"github.com/adnvilla/go_challenges/investment_api/src/pkg/gorm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ya que se garantiza que son multiplos de 100, entonces podemos inferir que los valores 300, 500, y 700 son los numeros primos de 3, 5 y 7
	//primes := []int{3, 5, 7}

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")

	// This can be moved to dispatcher
	db := gorm.GetConnection()
	s := service.NewCreditAssignmentService()
	repository := infrastructure.NewCreditAssignmentRepository(db)
	usecaseCreditAssignment := use_cases.NewCreateCreditAssignmentUseCase(s, repository)
	usecaseStatistics := use_cases.NewGetStatisticsUseCase(repository)
	handlerCreditAssignment := controllers.NewCreateCreditAssignmentHandler(usecaseCreditAssignment)
	handlerStatistics := controllers.NewStatisticsHandlerHandler(usecaseStatistics)

	api.POST("/credit-assignment", handlerCreditAssignment.CreateCreditAssignment)

	// This can be GET
	api.POST("/statistics", handlerStatistics.GetStatistics)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	// Run gin routing
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println("server failed on ListenAndServe")
		log.Println(err)
	}
}
