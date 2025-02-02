package models

type Basket struct {
	ID         int64 `json:"id"`           // bigserial в базе данных
	AuthUserID int   `json:"auth_user_id"` // integer в базе данных
	PositionID int   `json:"position_id"`  // integer в базе данных
}
