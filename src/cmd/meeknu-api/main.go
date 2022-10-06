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
		menu := menu.Menu{1, "this is my name", "my description", time.Now()}
		return c.JSON(http.StatusOK, menu)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
