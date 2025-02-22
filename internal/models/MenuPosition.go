package models

type MenuPosition struct {
	ID         int64 `json:"id"`
	MenuID     int   `json:"menu_id"`
	PositionID int   `json:"position_id"`
}
