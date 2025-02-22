package models

type Menu struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	AuthUserID int    `json:"auth_user_id"`
}
