package models

type AuthGroupPermissions struct {
	ID           int64 `json:"id"`            // bigserial в базе данных
	GroupID      int   `json:"group_id"`      // integer в базе данных
	PermissionID int   `json:"permission_id"` // integer в базе данных
}
