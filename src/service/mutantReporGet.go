package service

import (
	"main/src/repository"

	"github.com/gin-gonic/gin"
)

func ReportGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cantMutants, errMutants = repository.GetCantItemsByIsMutant(1)
		if errMutants != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error.",
			})
		} else {
			var cantHumans, errHumans = repository.GetCantItemsByIsMutant(0)
			if errHumans != nil {
				c.JSON(500, gin.H{
					"message": "Internal Server Error.",
				})
			} else {
				if cantMutants != 0 && cantHumans != 0 {
					c.JSON(200, gin.H{
						"count_mutant_dna": cantMutants,
						"count_human_dna":  cantHumans,
						"ratio":            calculateRatio(cantMutants, cantHumans),
					})
				} else {
					c.JSON(406, gin.H{
						"count_mutant_dna": cantMutants,
						"count_human_dna":  cantHumans,
						"ratio":            "No se puede calcular el ratio debido a que uno de los valores es igual a 0",
					})
				}
			}
		}
	}
}

//method used to calculate the ratio between 2 values
func calculateRatio(cantMutants int, cantHumans int) float64 {
	var greatestCommonDivisor = getGreatestCommonDivisor(cantMutants, cantHumans)
	return float64(float64(cantMutants/greatestCommonDivisor) / float64(cantHumans/greatestCommonDivisor))
}

//method used to calculate greatest common divisor of 2 values
func getGreatestCommonDivisor(num1 int, num2 int) int {
	if num1 == 0 {
		return num2
	}
	if num2 == 0 {
		return num1
	}
	return getGreatestCommonDivisor(num2%num1, num1)
}
