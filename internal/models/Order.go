package models

import "time"

type Order struct {
	ID                int       `json:"id"`
	TotalPrice        float64   `json:"total_price"`
	Status            string    `json:"status"`
	DeliveryAddress   string    `json:"delivery_address"`
	CreatedAt         time.Time `json:"created_at"`
	ClosedAt          time.Time `json:"closed_at"`
	AuthUserID        int       `json:"auth_user_id"`
	TrackNumber       string    `json:"track_number"`
	DelivererAccepted bool      `json:"deliverer_accepted"`
}
