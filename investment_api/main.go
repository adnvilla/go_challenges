package main

import (
	"github.com/adnvilla/go_challenges/investment_api/src/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ya que se garantiza que son multiplos de 100, entonces podemos inferir que los valores 300, 500, y 700 son los numeros primos de 3, 5 y 7
	//primes := []int{3, 5, 7}

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/credit-assignment", controllers.CreateCreditAssignment)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
