package service

import (
	"haproxy/handler"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.GET("/", handler.Handle)

	na := 2

	if len(os.Args) != na {
		log.Fatal("enter the port")
	}

	p := os.Args[1]

	e.Logger.Fatal(e.Start(":" + p))
}
