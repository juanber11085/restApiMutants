package main

import (
	"log"
	"main/src/repository"
	"main/src/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv() //we load the environment variables
	r := gin.Default()
	r.POST("/mutant", service.MutantValidatePost()) //we expose post service to validate if it is mutant according to dna
	r.GET("/stats", service.ReportGet())            //we expose get service to obtain the validated adns report
	r.Run()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	errCreateTable := repository.CreateTableIfNotExists()
	if errCreateTable != nil {
		log.Fatal("Error creando la tabla mutants")
	}
}
