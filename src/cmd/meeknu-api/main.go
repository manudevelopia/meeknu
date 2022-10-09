package main

import (
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		m := menu.Menu{ID: 1, Name: "this is my name", Description: "my description", CreatedOn: time.Now()}
		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
