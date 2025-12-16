package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gonzalo-wi/DiferenciasNautom/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	r := gin.Default()
	r.GET("api/differences/summary", handlers.GetDifferences)
	r.Run(":8080")
}
