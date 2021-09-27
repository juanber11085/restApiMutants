package main

import (
	"log"
	"main/src/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	r := gin.Default()
	r.POST("/mutant", service.MutantValidatePost())
	r.GET("/stats", service.ReportGet())
	r.Run()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
