package main

import (
	"fmt"
	"os"

	"github.com/GabrielVilarino/arch-hexagonal/adapter/input/routes"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/env"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start Application")

	err := godotenv.Overload()
	if err != nil {
		logger.Error("Error trying to load env", err)
		return
	}

	router := gin.Default()

	routes.InitRoutes(router)

	port := env.GetPort()
	fmt.Println(os.Getenv("NEWS_API_KEY"))
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Error(fmt.Sprintf("Error trying to start server on port: %s", port), err)
		return
	}
}
