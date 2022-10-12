package data

import (
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
)

type MenuRepository struct {
	Data *Data
}

func (ur *MenuRepository) GetAll() []menu.Menu {
	const query = "SELECT m_id, m_name, m_created_on FROM m_menu"
	rows, _ := ur.Data.DB.Query(query)
	var menus []menu.Menu
	for rows.Next() {
		var u menu.Menu
		_ = rows.Scan(&u.ID, &u.Name, &u.CreatedOn)
		menus = append(menus, u)
	}
	return menus
}
