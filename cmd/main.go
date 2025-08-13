package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"pratyutpannamati/db"
	"pratyutpannamati/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}

	database := db.InitDB()
	defer database.Close()

	r := gin.Default()
	routes.SetupRoutes(r, database)

	log.Println("🚀 Server running on http://localhost:8000")
	r.Run(":8000")
}
