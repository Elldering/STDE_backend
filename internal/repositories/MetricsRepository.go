package repositories

import (
	"STDE_proj/utils/db"
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

func GetUserCount() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM auth_user").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

func GetActiveProjectsCount() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM projects WHERE status = 'active'").Scan(&count)
	return count, err
}

func GetAverageTaskTime() (float64, error) {
	var avgTime float64
	err := db.DB.QueryRow("SELECT AVG(EXTRACT(EPOCH FROM (completed_at - created_at))) FROM tasks WHERE completed_at IS NOT NULL").Scan(&avgTime)
	return avgTime, err
}
