package models

type UserDocument struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	IsAccepted bool   `json:"is_accepted"`
	Type       string `json:"type"`
}
