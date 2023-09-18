package main

import (
	"gorest-10/configs"
	"gorest-10/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Failed load file env")
	}
}
