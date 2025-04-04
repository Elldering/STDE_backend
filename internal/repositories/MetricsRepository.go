package repositories

import (
	"STDE_proj/utils/database"
	"log"

	_ "github.com/lib/pq"
)

// Количество пользователей
func GetUserCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM auth_user").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Количество заблокированных пользователей
func GetBlockedUserCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM auth_user WHERE is_blocked = true").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Количество заказов
func GetOrderCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM \"order\"").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Количество активных заказов
func GetActiveOrderCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM \"order\" WHERE status = 'Создан'").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Средняя сумма заказа
func GetAverageOrderPrice() (float64, error) {
	var avgPrice float64
	err := database.DB.QueryRow("SELECT AVG(total_price) FROM \"order\" WHERE total_price IS NOT NULL").Scan(&avgPrice)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return avgPrice, nil
}

// Общая выручка
func GetTotalRevenue() (float64, error) {
	var total float64
	err := database.DB.QueryRow("SELECT SUM(total_price) FROM \"order\" WHERE total_price IS NOT NULL").Scan(&total)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return total, nil
}

// Количество позиций в заказах
func GetOrderPositionCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM order_position").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Среднее количество товаров в одном заказе
func GetAverageOrderItems() (float64, error) {
	var avgItems float64
	err := database.DB.QueryRow("SELECT AVG(item_count) FROM (SELECT order_id, COUNT(*) as item_count FROM order_position GROUP BY order_id) as order_items").Scan(&avgItems)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return avgItems, nil
}

// Количество отзывов
func GetReviewCount() (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM reviews").Scan(&count)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return count, nil
}

// Средний рейтинг отзывов
func GetAverageReviewRating() (float64, error) {
	var avgRating float64
	err := database.DB.QueryRow("SELECT AVG(grade) FROM reviews").Scan(&avgRating)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v", err)
		return 0, err
	}
	return avgRating, nil
}
