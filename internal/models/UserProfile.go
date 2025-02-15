package models

type UserProfile struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MiddleName   string  `json:"middle_name"`
	Rating       float32 `json:"rating"`
	ProfileImage string  `json:"profile_image"`
}
