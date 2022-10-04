package menu

import "time"

type Menu struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}
