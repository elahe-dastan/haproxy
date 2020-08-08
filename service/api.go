package service

import (
	"github.com/labstack/echo/v4"
	"haproxy/handler"
	"log"
	"os"
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
