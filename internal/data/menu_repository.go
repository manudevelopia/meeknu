package data

import (
	"context"
	"github.com/manudevelopia/meeknu-api/src/pkg/menu"
	"log"
)

type MenuRepository struct {
	Data *Data
}

func (ur *MenuRepository) GetAll() []menu.Menu {
	const query = "SELECT m_id, m_name, m_created_on FROM m_menu"
	rows, err := ur.Data.DB.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var menus []menu.Menu
	for rows.Next() {
		var u menu.Menu
		_ = rows.Scan(&u.ID, &u.Name, &u.CreatedOn)
		menus = append(menus, u)
	}
	return menus
}
