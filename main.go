package main

import (
	"fmt"
	"log"
	"os"
	"test-echo/database"
	"test-echo/routes"

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
	defer db.Close()

	e := echo.New()
	v1 := e.Group("v1")
	r := routes.NewRoute(db)
	r.Register(v1)

	appPort := os.Getenv("APP_PORT")
	serverAddress := fmt.Sprintf(":%s", appPort)

	e.Logger.Fatal(e.Start(serverAddress))
}
