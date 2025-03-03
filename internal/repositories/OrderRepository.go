package repositories

//
//import (
//	"STDE_proj/internal/models"
//	"STDE_proj/utils/database"
//	"database/sql"
//	"fmt"
//	"log"
//)
//
//func GetOrder() ([]models.Order, error) {
//
//	if database.DB == nil {
//		log.Println("Ошибка: подключение к базе данных не инициализировано")
//		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
//	}
//
//	query, err := database.DB.Query("SELECT id, total_price, status, delivery_address, create_at, closed_at, auth_user_id, track_number, deliverer_accepted FROM Order")
//
//	if err != nil {
//		return nil, err
//	}
//	defer query.Close()
//
//	var orders []models.Order
//	for query.Next() {
//		var order models.Order
//		var closed_at sql.NullString
//		if err := query.Scan(&review.ID, &review.AuthUserSenderID, &review.AuthUserRecipientID, &review.Grade, &comment, &review.CreatedAt); err != nil {
//			return nil, err
//		}
//		review.Comment = comment.String
//		reviews = append(reviews, review)
//	}
//	return reviews, nil
//}
