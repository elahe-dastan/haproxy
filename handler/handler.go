package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func Handle(c echo.Context) error {
	sd := 5
	time.Sleep(time.Duration(sd) * time.Second)

	return c.JSON(http.StatusOK, "hi "+os.Args[1])
}
