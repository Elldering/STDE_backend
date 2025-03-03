package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetMenu() ([]models.Menu, error) {
	rows, err := database.DB.Query("SELECT id, name, auth_user_id FROM menu")
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var menus []models.Menu
	for rows.Next() {
		var menu models.Menu
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.AuthUserID); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func GetMenuById(id int) (models.Menu, error) {
	row := database.DB.QueryRow("SELECT id, name, auth_user_id FROM menu WHERE id = $1", id)
	var menu models.Menu
	if err := row.Scan(&menu.ID, &menu.Name, &menu.AuthUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Menu{}, fmt.Errorf("меню с id %d не найдено", id)
		}
		return models.Menu{}, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	return menu, nil
}

func PostMenu(menu models.Menu) error {
	query := "INSERT INTO menu (name, auth_user_id) VALUES ($1, $2)"
	result, err := database.DB.Exec(query, menu.Name, menu.AuthUserID)
	if err != nil {
		log.Printf("Ошибка при добавлении меню: %v", err)
		return fmt.Errorf("ошибка при добавлении меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("меню не было добавлено")
	}
	return nil
}

func PutMenu(id int, menu models.Menu) error {
	query := "UPDATE menu SET name = $1, auth_user_id = $2 WHERE id = $3"
	result, err := database.DB.Exec(query, menu.Name, menu.AuthUserID, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("меню с id %d не найдено", id)
	}
	return nil
}

func DeleteMenu(id int) error {
	query := "DELETE FROM menu WHERE id = $1"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении меню: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("меню с id %d не найдено", id)
	}
	return nil
}
