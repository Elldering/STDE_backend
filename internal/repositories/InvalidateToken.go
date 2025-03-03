package repositories

import (
	"STDE_proj/utils/database"
	"time"
)

// InvalidateToken добавляет токен в таблицу invalid_tokens
func InvalidateToken(access string, refresh string) error {
	access = access[len("Bearer "):]
	refresh = refresh[len("Bearer "):]
	query := "INSERT INTO invalid_tokens (access, refresh, created_at) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, access, refresh, time.Now())
	return err
}

// IsTokenInvalidated проверяет, находится ли токен в таблице invalid_tokens
func IsAccessTokenInvalidated(access string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM invalid_tokens WHERE access = $1"
	err := database.DB.QueryRow(query, access).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func IsRefreshTokenInvalidated(refresh string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM invalid_tokens WHERE refresh = $1"
	err := database.DB.QueryRow(query, refresh).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
