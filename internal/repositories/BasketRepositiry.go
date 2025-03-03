package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetBasket() ([]models.Basket, error) {
	rows, err := database.DB.Query("SELECT id, auth_user_id, position_id FROM basket")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var baskets []models.Basket
	for rows.Next() {
		var basket models.Basket
		if err := rows.Scan(&basket.ID, &basket.AuthUserID, &basket.PositionID); err != nil {
			return nil, err
		}
		baskets = append(baskets, basket)
	}
	return baskets, nil
}
func GetBasketById(id int) (models.Basket, error) {
	row := database.DB.QueryRow("SELECT id, auth_user_id, position_id FROM basket where id=$1", id)
	var basket models.Basket
	if err := row.Scan(&basket.ID, &basket.AuthUserID, &basket.PositionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Basket{}, fmt.Errorf("блюдо с id %d: %w не найдено", id, err)
		}
		return models.Basket{}, err
	}
	return basket, nil
}
func PostBasket(agp models.Basket) error {
	

	query := "INSERT INTO basket (auth_user_id, position_id) VALUES ($1, $2)"
	result, err := database.DB.Exec(query, agp.AuthUserID, agp.PositionID)
	if err != nil {
		log.Printf("Ошибка при добавлении записи в корзину: %v", err)
		return fmt.Errorf("ошибка при добавлении записи в корзину: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при получении количества затронутых строк: %v", err)
		return fmt.Errorf("ошибка при получении количества затронутых строк: %v", err)
	}

	if rowsAffected == 0 {
		log.Println("Предупреждение: запись не была добавлена в корзину")
		return errors.New("запись не была добавлена в корзину")
	}

	return nil

}

func DeleteBasketPosition(ctx context.Context, id int) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}
	_, err := database.DB.Exec("DELETE FROM basket WHERE id = $1", id)
	return err
}
func DeleteBasketByUserID(ctx context.Context, id int) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}
	_, err := database.DB.Exec("DELETE FROM basket WHERE auth_user_id = $1", id)
	return err
}
