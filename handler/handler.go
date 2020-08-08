package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

func Handle(c echo.Context) error {
	time.Sleep(5 * time.Second)

	return c.JSON(http.StatusOK, "hi " + os.Args[1])
}
