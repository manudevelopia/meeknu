package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
	"net/http"
)

func main() {
	e := echo.New()
	connStr := ""
	db, _ := sql.Open("postgres", connStr)

	e.GET("/", func(c echo.Context) error {
		rows, _ := db.Query("SELECT m_id, m_name, m_created_on FROM m_menu")
		var menus []menu.Menu
		for rows.Next() {
			var u menu.Menu
			_ = rows.Scan(&u.ID, &u.Name, &u.CreatedOn)
			menus = append(menus, u)
		}
		return c.JSON(http.StatusOK, menus)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
