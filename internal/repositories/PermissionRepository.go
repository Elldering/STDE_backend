package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

func GetAllPermissions() ([]models.Permission, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, codename, description FROM permission")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		var permission models.Permission
		if err := rows.Scan(&permission.ID, &permission.Codename, &permission.Description); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func GetPermissionById(id int) (models.Permission, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.Permission{}, errors.New("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, codename, description FROM permission WHERE id=$1", id)

	var permission models.Permission
	if err := row.Scan(&permission.ID, &permission.Codename, &permission.Description); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Permission с id %d не найдена", id)
			return models.Permission{}, fmt.Errorf("Permission с id %d не найдена", id)
		}
		log.Printf("Ошибка при получении Permission с id %d: %v", id, err)
		return models.Permission{}, fmt.Errorf("ошибка при получении Permission: %v", err)
	}

	return permission, nil
}

func PostPermission(agp models.Permission) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	if strings.TrimSpace(agp.Codename) == "" || strings.TrimSpace(agp.Description) == "" {
		log.Println("Ошибка: пустые поля codename или description")
		return errors.New("поля codename и description не могут быть пустыми")
	}

	query := "INSERT INTO permission (codename, description) VALUES ($1, $2)"
	result, err := db.DB.Exec(query, agp.Codename, agp.Description)
	if err != nil {
		log.Printf("Ошибка при добавлении Permission: %v", err)
		return fmt.Errorf("ошибка при добавлении Permission: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества добавленных строк: %v", err)
		return fmt.Errorf("ошибка при проверке количества добавленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Println("Ошибка: Permission не была добавлена")
		return errors.New("Permission не была добавлена")
	}

	return nil
}

func DeletePermission(id int) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	result, err := db.DB.Exec("DELETE FROM permission WHERE id=$1", id)
	if err != nil {
		log.Printf("Ошибка при удалении Permission с id %d: %v", id, err)
		return fmt.Errorf("ошибка при удалении Permission: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества удаленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: Permission с id %d не найдена", id)
		return fmt.Errorf("Permission с id %d не найдена", id)
	}

	return nil
}

func PutPermission(id int, agp models.Permission) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	if strings.TrimSpace(agp.Codename) == "" || strings.TrimSpace(agp.Description) == "" {
		log.Println("Ошибка: пустые поля codename или description")
		return errors.New("поля codename и description не могут быть пустыми")
	}

	row := db.DB.QueryRow("SELECT id FROM permission WHERE id=$1", id)
	var existingID int
	if err := row.Scan(&existingID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Permission с id %d не найдена", id)
			return fmt.Errorf("Permission с id %d не найдена", id)
		}
		log.Printf("Ошибка при проверке существования Permission с id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке существования Permission: %v", err)
	}

	result, err := db.DB.Exec("UPDATE permission SET codename=$1, description=$2 WHERE id=$3", agp.Codename, agp.Description, id)
	if err != nil {
		log.Printf("Ошибка при обновлении Permission с id %d: %v", id, err)
		return fmt.Errorf("ошибка при обновлении Permission: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества обновленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: Permission с id %d не найдена", id)
		return fmt.Errorf("Permission с id %d не найдена", id)
	}

	return nil
}
