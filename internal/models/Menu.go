package models

type Menu struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PositionID int    `json:"position_id"`
	AuthUserID int    `json:"auth_user_id"`
}
