package data

import (
	"github.com/jackc/pgx/v5"
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
)

type MenuRepository struct {
	Data *Data
}

func (ur *MenuRepository) GetAll() []menu.Menu {
	const query = "SELECT m_id, m_name, m_created_on FROM m_menu"
	rows := ur.Data.Query(query)
	return menuMapper(rows)
}

func menuMapper(rows pgx.Rows) []menu.Menu {
	var menus []menu.Menu
	for rows.Next() {
		var u menu.Menu
		_ = rows.Scan(&u.ID, &u.Name, &u.CreatedOn)
		menus = append(menus, u)
	}
	return menus
}
