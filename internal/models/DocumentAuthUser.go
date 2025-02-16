package models

type DocumentAuthUser struct {
	ID             int64 `json:"id"`               // bigserial в базе данных
	AuthUserID     int   `json:"auth_user_id"`     // integer в базе данных
	UserDocumentID int   `json:"user_document_id"` // integer в базе данных
}
