package models

type OrderPosition struct {
	ID         int64 `json:"id"`
	OrderID    int   `json:"order_id"`
	PositionID int   `json:"position_id"`
}
