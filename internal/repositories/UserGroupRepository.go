package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"log"
)

func GetAllUserGroups() ([]models.UserGroup, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, name FROM user_group")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userGroups []models.UserGroup
	for rows.Next() {
		var userGroup models.UserGroup
		if err := rows.Scan(&userGroup.ID, &userGroup.Name); err != nil {
			return nil, err
		}
		userGroups = append(userGroups, userGroup)
	}

	return userGroups, nil
}

func GetUserGroupById(id int) (models.UserGroup, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.UserGroup{}, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, name FROM user_group WHERE id=$1", id)

	var userGroup models.UserGroup
	if err := row.Scan(&userGroup.ID, &userGroup.Name); err != nil {
		if err == sql.ErrNoRows {
			return models.UserGroup{}, fmt.Errorf("группа пользователей с id %d не найдена", id)
		}
		return models.UserGroup{}, err
	}

	return userGroup, nil
}

func PostUserGroup(agp models.UserGroup) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO user_group ( name) VALUES ($1)"
	_, err := db.DB.Exec(query, agp.Name)
	return err
}

func DeleteUserGroup(id int) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Удаление группы пользователей
	result, err := db.DB.Exec("DELETE FROM user_group WHERE id=$1", id)
	if err != nil {
		log.Println("Ошибка при удалении группы пользователей:", err)
		return fmt.Errorf("ошибка при удалении группы пользователей: %v", err)
	}

	// Проверка, была ли удалена хотя бы одна строка
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества удаленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("группа пользователей с id %d не найдена", id)
	}

	return nil
}

func PutUserGroup(id int, agp models.UserGroup) error {
	// Проверка инициализации подключения к базе данных
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Проверка существования группы пользователей
	row := db.DB.QueryRow("SELECT id, name FROM user_group WHERE id=$1", id)
	var userGroup models.UserGroup
	if err := row.Scan(&userGroup.ID, &userGroup.Name); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("группа пользователей с id %d не найдена", id)
		}
		return err
	}

	// Обновление данных группы пользователей
	result, err := db.DB.Exec("UPDATE user_group SET name=$1 WHERE id=$2", agp.Name, id)
	if err != nil {
		log.Println("Ошибка при обновлении группы пользователей:", err)
		return fmt.Errorf("ошибка при обновлении группы пользователей: %v", err)
	}

	// Проверка, была ли обновлена хотя бы одна строка
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества обновленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("группа пользователей с id %d не найдена", id)
	}

	return nil
}
