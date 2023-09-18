package main

import (
	"gorest-10/configs"
	"gorest-10/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(configs.GetPort())
}

// func loadEnv() {
// 	if err := godotenv.Load(); err != nil {
// 		panic("Failed load file env")
// 	}
// }
