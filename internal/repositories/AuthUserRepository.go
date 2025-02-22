package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"errors"
	"fmt"
	"log"
)

//func GetAllAuthUser()([]models.AuthUser, error){
//	if database.DB == nil {
//		log.Println("Ошибка: подключение к базе данных не инициализировано")
//		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
//	}
//	rows, err := database.DB.Query("SELECT id, password, last_login, email, phone_number, is_super_user, is_email, is_number_verify, is_blocked, FROM permission")
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var permissions []models.Permission
//	for rows.Next() {
//		var permission models.Permission
//		if err := rows.Scan(&permission.ID, &permission.Codename, &permission.Description); err != nil {
//			return nil, err
//		}
//		permissions = append(permissions, permission)
//	}
//
//	return permissions, nil
//}

func PostAuthUser(data models.AuthUser) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO auth_user (password, email, phone_number) VALUES ($1, $2, $3)"
	result, err := database.DB.Exec(query, data.Password, data.Email, data.PhoneNumber)
	if err != nil {
		log.Printf("Ошибка при добавлении: %v", err)
		return fmt.Errorf("ошибка при добавлении: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества добавленных строк: %v", err)
		return fmt.Errorf("ошибка при проверке количества добавленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Println("Ошибка: пользователь не был добавлен")
		return errors.New("пользователь не был добавлен")
	}

	return nil
}

func DeleteAuthUser(id int) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	result, err := database.DB.Exec("DELETE FROM auth_user WHERE id=$1", id)
	if err != nil {
		log.Printf("Ошибка при удалении Пользователя с id %d: %v", id, err)
		return fmt.Errorf("ошибка при удалении Пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества удаленных строк для id %d: %v", id, err)
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: Пользователь с id %d не найдена", id)
		return fmt.Errorf("пользователь с id %d не найдена", id)
	}

	return nil
}
