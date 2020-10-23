package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	if c.Request().Method == "GET" {



		return c.Render(http.StatusOK, "index.html", "")

	} else {

		return c.String(404,"404")
	}
}
