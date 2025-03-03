package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetMenuPositions() ([]models.MenuPosition, error) {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, errors.New("подключение к базе данных не инициализировано")
	}

	rows, err := database.DB.Query("SELECT id, menu_id, position_id FROM menu_position")
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var menuPositions []models.MenuPosition
	for rows.Next() {
		var mp models.MenuPosition
		if err := rows.Scan(&mp.ID, &mp.MenuID, &mp.PositionID); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}
		menuPositions = append(menuPositions, mp)
	}
	return menuPositions, nil
}

func GetMenuPositionById(id int64) (models.MenuPosition, error) {
	row := database.DB.QueryRow("SELECT id, menu_id, position_id FROM menu_position WHERE id = $1", id)
	var mp models.MenuPosition
	if err := row.Scan(&mp.ID, &mp.MenuID, &mp.PositionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.MenuPosition{}, fmt.Errorf("позиция меню с id %d не найдена", id)
		}
		return models.MenuPosition{}, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	return mp, nil
}

func PostMenuPosition(mp models.MenuPosition) error {

	query := "INSERT INTO menu_position (menu_id, position_id) VALUES ($1, $2)"
	result, err := database.DB.Exec(query, mp.MenuID, mp.PositionID)
	if err != nil {
		log.Printf("Ошибка при добавлении позиции меню: %v", err)
		return fmt.Errorf("ошибка при добавлении позиции меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("позиция меню не была добавлена")
	}
	return nil
}

func PutMenuPosition(id int64, mp models.MenuPosition) error {
	query := "UPDATE menu_position SET menu_id = $1, position_id = $2 WHERE id = $3"
	result, err := database.DB.Exec(query, mp.MenuID, mp.PositionID, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении позиции меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("позиция меню с id %d не найдена", id)
	}
	return nil
}

func DeleteMenuPosition(id int64) error {

	query := "DELETE FROM menu_position WHERE id = $1"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении позиции меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("позиция меню с id %d не найдена", id)
	}
	return nil
}
