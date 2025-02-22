package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetAuthGroupPermissions() ([]models.AuthGroupPermissions, error) {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	rows, err := database.DB.Query("SELECT id, group_id, permission_id FROM auth_group_permissions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var AuthGroupPermissions []models.AuthGroupPermissions
	for rows.Next() {
		var AuthGroupPermission models.AuthGroupPermissions
		if err := rows.Scan(&AuthGroupPermission.ID, &AuthGroupPermission.GroupID, &AuthGroupPermission.PermissionID); err != nil {
			return nil, err
		}
		AuthGroupPermissions = append(AuthGroupPermissions, AuthGroupPermission)
	}

	return AuthGroupPermissions, nil
}

func GetAuthGroupPermissionsId(id int) (models.AuthGroupPermissions, error) {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.AuthGroupPermissions{}, fmt.Errorf("подключение к базе данных не инициализировано")
	}
	row := database.DB.QueryRow("SELECT id, group_id, permission_id FROM auth_group_permissions WHERE id=$1", id)
	var data models.AuthGroupPermissions
	if err := row.Scan(&data.ID, &data.GroupID, &data.PermissionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Связь групп и прав доступа с id %d не найдена", id)
			return models.AuthGroupPermissions{}, fmt.Errorf("связь групп и прав доступа с id %d не найдена", id)
		}
		log.Printf("Ошибка при получении связи групп и прав доступа с id %d: %v", id, err)
		return models.AuthGroupPermissions{}, fmt.Errorf("ошибка при получении связи групп и прав доступа: %v", err)
	}
	return data, nil
}

func PutAuthGroupPermissions(id int, data models.AuthGroupPermissions) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}
	if data.GroupID == 0 || data.PermissionID == 0 {
		log.Println("Ошибка: пустые поля group_id или permission_id")
		return errors.New("поля group_id и permission_id не могут быть пустыми")
	}

	row := database.DB.QueryRow("SELECT id FROM auth_group_permissions WHERE id=$1", id)
	var existingID int
	if err := row.Scan(&existingID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Ошибка: Связь группы с правами доступа с id %d не найдена", id)
			return fmt.Errorf("связь группы с правами доступа с id %d не найдена", id)
		}
		log.Printf("Ошибка при проверке существования связи группы с правами доступа с id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке существования связи группы с правами доступа: %v", err)
	}

	exec, err := database.DB.Exec("UPDATE auth_group_permissions SET group_id=$1, permission_id=$2 WHERE id=$3", data.GroupID, data.PermissionID, id)
	if err != nil {
		log.Printf("Ошибка при обновлении связи группы с правами доступа с id %d: %v", id, err)
		return fmt.Errorf("ошибка при обновлении связи группы с правами доступа: %v", err)
	}

	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества обновленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: связи группы с правами доступа с id %d не найдена", id)
		return fmt.Errorf("связи группы с правами доступа с id %d не найдена", id)
	}

	return nil
}

func DeleteAuthGroupPermissions(id int) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Удаление группы пользователей
	exec, err := database.DB.Exec("DELETE FROM auth_group_permissions WHERE id=$1", id)
	if err != nil {
		log.Println("Ошибка при удалении связи группы с правами доступа:", err)
		return fmt.Errorf("ошибка при удалении связи группы с правами доступа: %v", err)
	}

	// Проверка, была ли удалена хотя бы одна строка
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		log.Println("Ошибка при проверке количества удаленных строк:", err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("связи группы с правами доступа с id %d не найдена", id)
	}

	return nil
}

// Функция для создания связи группы и прав доступа
func PostAuthGroupPermission(agp models.AuthGroupPermissions) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO auth_group_permissions (group_id, permission_id) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, agp.GroupID, agp.PermissionID)
	return err
}
