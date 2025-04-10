package main

import (
	"portfolio-backend/config"
	"portfolio-backend/database"
	"portfolio-backend/middlewares"
	"portfolio-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.SetupLogger()
	database.Connect()

	// Buat instance router di sini
	r := gin.Default()

	// Pasang CORS di sini (PASTI BERFUNGSI)
	r.Use(middlewares.CORSMiddleware())

	// Kirim router ke fungsi setup
	routes.SetupRouter(r)

	// Run
	r.Run(":" + config.EnvPort())
}
