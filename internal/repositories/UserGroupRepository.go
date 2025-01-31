package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"fmt"
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
