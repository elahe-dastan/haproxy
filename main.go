package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/", handler)

	if len(os.Args) < 2 {
		log.Fatal("enter the port")
	}

	p := os.Args[1]

	e.Logger.Fatal(e.Start(":" + p))
}

func handler(c echo.Context) error {
	time.Sleep(5 * time.Second)

	return c.JSON(http.StatusOK, "hi")
}