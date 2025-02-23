package models

import "time"

type AuthUser struct {
	ID             int       `json:"id"`
	Password       string    `json:"password"`
	LastLogin      time.Time `json:"last_login"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	IsSuperUser    bool      `json:"is_super_user"`
	IsEmail        bool      `json:"is_email"`
	IsNumberVerify bool      `json:"is_number_verify"`
	IsBlocked      bool      `json:"is_blocked"`
	IsFrozen       bool      `json:"is_frozen"`
	DateJoined     time.Time `json:"date_joined"`
	UserProfileID  int       `json:"user_profile_id"`
	Balance        float64   `json:"balance"`
	UserRID        int       `json:"user_r_id"`
	Login          string    `json:"login"`
	TypeLogin      string    `json:"type_login"`
}
