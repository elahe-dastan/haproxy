package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/", handler)
	
}

func handler(c echo.Context) error {
	time.Sleep(5 * time.Second)
	fmt.Println("hi")

	return nil
}