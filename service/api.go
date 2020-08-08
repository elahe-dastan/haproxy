package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func Run() {
	e := echo.New()

	e.GET("/", handler)

	if len(os.Args) < 2 {
		log.Fatal("enter the port")
	}

	p := os.Args[1]

	e.Logger.Fatal(e.Start(":" + p))
}
