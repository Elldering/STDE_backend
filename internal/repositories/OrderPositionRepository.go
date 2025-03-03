package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetOrderPositions() ([]models.OrderPosition, error) {
	rows, err := database.DB.Query("SELECT id, order_id, position_id FROM order_position")
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var orderPositions []models.OrderPosition
	for rows.Next() {
		var op models.OrderPosition
		if err := rows.Scan(&op.ID, &op.OrderID, &op.PositionID); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}
		orderPositions = append(orderPositions, op)
	}
	return orderPositions, nil
}

func GetOrderPositionById(id int64) (models.OrderPosition, error) {
	row := database.DB.QueryRow("SELECT id, order_id, position_id FROM order_position WHERE id = $1", id)
	var op models.OrderPosition
	if err := row.Scan(&op.ID, &op.OrderID, &op.PositionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.OrderPosition{}, fmt.Errorf("позиция заказа с id %d не найдена", id)
		}
		return models.OrderPosition{}, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	return op, nil
}

func PostOrderPosition(op models.OrderPosition) error {
	query := "INSERT INTO order_position (order_id, position_id) VALUES ($1, $2)"
	result, err := database.DB.Exec(query, op.OrderID, op.PositionID)
	if err != nil {
		log.Printf("Ошибка при добавлении позиции заказа: %v", err)
		return fmt.Errorf("ошибка при добавлении позиции заказа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("позиция заказа не была добавлена")
	}
	return nil
}

func PutOrderPosition(id int64, op models.OrderPosition) error {
	query := "UPDATE order_position SET order_id = $1, position_id = $2 WHERE id = $3"
	result, err := database.DB.Exec(query, op.OrderID, op.PositionID, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении позиции заказа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("позиция заказа с id %d не найдена", id)
	}
	return nil
}

func DeleteOrderPosition(id int64) error {
	query := "DELETE FROM order_position WHERE id = $1"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении позиции заказа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("позиция заказа с id %d не найдена", id)
	}
	return nil
}
