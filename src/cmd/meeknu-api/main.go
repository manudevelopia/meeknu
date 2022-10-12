package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/manudevelopia/meeknu-api/internal/data"
	"log"
	"net/http"
)

func main() {
	server := echo.New()
	menuRepository := &data.MenuRepository{Data: data.New()}

	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, menuRepository.GetAll())
	})

	log.Fatal(server.Start(":1323"))
}
