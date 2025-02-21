package repositories

import (
	"STDE_proj/utils/db"
	"time"
)

// InvalidateToken добавляет токен в таблицу invalid_tokens
func InvalidateToken(token string) error {
	tokenString := token[len("Bearer "):]
	query := "INSERT INTO invalid_tokens (token, created_at) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, tokenString, time.Now())
	return err
}

// IsTokenInvalidated проверяет, находится ли токен в таблице invalid_tokens
func IsTokenInvalidated(token string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM invalid_tokens WHERE token = $1"
	err := db.DB.QueryRow(query, token).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
