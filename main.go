package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yogapratama23/go-echo-boilerplate/database"
	"github.com/yogapratama23/go-echo-boilerplate/middlewares"
	"github.com/yogapratama23/go-echo-boilerplate/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	db := database.InitDB()

	e := echo.New()
	e.Use(middlewares.LoggerMiddleware())

	v1 := e.Group("v1")
	r := routes.NewRoute(db)
	r.RegisterV1(v1)

	appPort := os.Getenv("APP_PORT")
	serverAddress := fmt.Sprintf(":%s", appPort)

	e.Logger.Fatal(e.Start(serverAddress))
}
