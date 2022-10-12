package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/manudevelopia/meeknu-api/internal/data"
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
	"log"
	"net/http"
)

func main() {
	d := data.New()
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		rows, _ := d.DB.Query("SELECT m_id, m_name, m_created_on FROM m_menu")
		var menus []menu.Menu
		for rows.Next() {
			var u menu.Menu
			_ = rows.Scan(&u.ID, &u.Name, &u.CreatedOn)
			menus = append(menus, u)
		}
		return c.JSON(http.StatusOK, menus)
	})

	log.Fatal(server.Start(":1323"))
}
