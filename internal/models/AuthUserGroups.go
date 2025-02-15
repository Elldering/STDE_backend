package models

type AuthGroupGroups struct {
	ID      int64 `json:"id"`       // bigserial в базе данных
	GroupID int   `json:"group_id"` // integer в базе данных
	UserID  int   `json:"user_id"`  // integer в базе данных
}
