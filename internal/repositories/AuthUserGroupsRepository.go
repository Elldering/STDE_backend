package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetAuthUserGroupsAll() ([]models.AuthUserGroups, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}
	query, err := db.DB.Query("SELECT id, group_id, user_id FROM auth_user_groups")
	if err != nil {
		return nil, err
	}
	defer query.Close()
	var authUserGroups []models.AuthUserGroups
	for query.Next() {
		var authUserGroup models.AuthUserGroups
		if err := query.Scan(&authUserGroup.ID, &authUserGroup.GroupID, &authUserGroup.UserID); err != nil {
			return nil, err
		}
		authUserGroups = append(authUserGroups, authUserGroup)

	}
	return authUserGroups, nil
}

func GetAuthUserGroupsById(id int) (models.AuthUserGroups, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.AuthUserGroups{}, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, group_id, user_id FROM auth_user_groups WHERE id = $1", id)
	var data models.AuthUserGroups
	if err := row.Scan(&data.ID, &data.GroupID, &data.UserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Связь групп и пользователя с id %d не найдена", id)
			return models.AuthUserGroups{}, fmt.Errorf("связь групп и пользователя с id %d не найдена", id)
		}
		log.Printf("Ошибка при получении связи групп и пользователя с id %d: %v", id, err)
		return models.AuthUserGroups{}, fmt.Errorf("ошибка при получении связи групп и пользователя: %v", err)
	}
	return data, nil
}

func PostAuthUserGroups(data models.AuthUserGroups) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO auth_user_groups (group_id, user_id) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, data.GroupID, data.UserID)
	return err
}

func PutAuthUserGroups(id int, data models.AuthUserGroups) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	if data.GroupID == 0 || data.UserID == 0 {
		log.Println("Ошибка: пустые поля group_id или user_id")
		return errors.New("поля group_id и user_id не могут быть пустыми")
	}

	row := db.DB.QueryRow("SELECT id FROM auth_user_groups WHERE id=$1", id)
	var existingID int
	if err := row.Scan(&existingID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Связь группы и пользователями с id %d не найдена", id)
			return fmt.Errorf("связь группы и пользователями с id %d не найдена", id)
		}
		log.Printf("Ошибка при проверке существования связи группы и пользователями с id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке существования связи группы и пользователями: %v", err)
	}

	exec, err := db.DB.Exec("UPDATE auth_user_groups SET group_id=$1, user_id=$2 WHERE id=$3", data.GroupID, data.UserID, id)
	if err != nil {
		log.Printf("Ошибка при обновлении связи группы и пользователями с id %d: %v", id, err)
		return fmt.Errorf("ошибка при обновлении связи группы и пользователями: %v", err)
	}

	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества обновленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: связи группы и пользователями с id %d не найдена", id)
		return fmt.Errorf("связи группы и пользователями с id %d не найдена", id)
	}

	return nil
}

func DeleteAuthUserGroups(id int) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Удаление группы пользователей
	exec, err := db.DB.Exec("DELETE FROM auth_user_groups WHERE id=$1", id)
	if err != nil {
		log.Println("Ошибка при удалении связи группы и пользователя:", err)
		return fmt.Errorf("ошибка при удалении связи группы и пользователя: %v", err)
	}

	// Проверка, была ли удалена хотя бы одна строка
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества удаленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("связи группы и пользователя с id %d не найдена", id)
	}

	return nil
}
