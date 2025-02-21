package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetUserCount() (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM auth_user").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}
