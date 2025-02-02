package models

import "time"

type Reviews struct {
	ID                  int       `json:"id"`
	AuthUserSenderID    int       `json:"auth_user_sender_id"`
	AuthUserRecipientID int       `json:"auth_user_recipient_id"`
	Grade               int16     `json:"grade"`
	Comment             string    `json:"comment"`
	CreatedAt           time.Time `json:"created_at"`
}
