package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetReviewsAll() ([]models.Reviews, error) {

	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query, err := database.DB.Query("SELECT id, auth_user_sender_id, auth_user_recipient_id, grade, comment, created_at FROM Reviews")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var reviews []models.Reviews
	for query.Next() {
		var review models.Reviews
		var comment sql.NullString
		if err := query.Scan(&review.ID, &review.AuthUserSenderID, &review.AuthUserRecipientID, &review.Grade, &comment, &review.CreatedAt); err != nil {
			return nil, err
		}
		review.Comment = comment.String
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func GetReviewsById(id int) (models.Reviews, error) {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.Reviews{}, fmt.Errorf("подключение к базе данных не инициализировано")
	}
	query := database.DB.QueryRow("SELECT id, auth_user_sender_id, auth_user_recipient_id, grade, comment, created_at FROM Reviews WHERE id=$1", id)
	var review models.Reviews
	if err := query.Scan(&review.ID, &review.AuthUserSenderID, &review.AuthUserRecipientID, &review.Grade, &review.Comment, &review.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Отзыв с id %d не найдена", id)
			return models.Reviews{}, fmt.Errorf("отзыв с id %d не найдена", id)
		}
		log.Printf("Ошибка при получении отзыва с id %d: %v", id, err)
		return models.Reviews{}, fmt.Errorf("ошибка при получении отзыва: %v", err)
	}
	return review, nil
}

func PostReviews(data models.Reviews) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO Reviews (auth_user_sender_id, auth_user_recipient_id, grade, comment ) VALUES ($1, $2, $3, $4)"
	_, err := database.DB.Exec(query, data.AuthUserSenderID, data.AuthUserRecipientID, data.Grade, data.Comment)
	return err
}
