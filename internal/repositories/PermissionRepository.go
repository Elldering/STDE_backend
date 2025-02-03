package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"fmt"
	"log"
)

func GetAllPermissions() ([]models.Permisson, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, codename, description FROM permisson")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissons []models.Permisson
	for rows.Next() {
		var permisson models.Permisson
		if err := rows.Scan(&permisson.ID, &permisson.Codename, &permisson.Description); err != nil {
			return nil, err
		}
		permissons = append(permissons, permisson)
	}

	return permissons, nil
}

func GetPermissionById(id int) (models.Permisson, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.Permisson{}, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, codename, description FROM permisson WHERE id=$1", id)

	var Permisson models.Permisson
	if err := row.Scan(&Permisson.ID, &Permisson.Codename, &Permisson.Description); err != nil {
		if err == sql.ErrNoRows {
			return models.Permisson{}, fmt.Errorf("группа пользователей с id %d не найдена", id)
		}
		return models.Permisson{}, err
	}

	return Permisson, nil
}

func PostPermission(agp models.Permisson) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO permisson ( codename, description) VALUES ($1,$2)"
	_, err := db.DB.Exec(query, agp.Codename, agp.Description)
	return err
}

func DeletePermission(id int) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Удаление группы пользователей
	result, err := db.DB.Exec("DELETE FROM permisson WHERE id=$1", id)
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

func PutPermission(id int, agp models.Permisson) error {
	// Проверка инициализации подключения к базе данных
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	// Проверка существования группы пользователей
	row := db.DB.QueryRow("SELECT id, codename, description FROM permisson WHERE id=$1", id)
	var Permisson models.Permisson
	if err := row.Scan(&Permisson.ID, &Permisson.Codename, &Permisson.Description); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("группа пользователей с id %d не найдена", id)
		}
		return err
	}

	// Обновление данных группы пользователей
	result, err := db.DB.Exec("UPDATE permisson SET codename=$1, description=$2 WHERE id=$3", agp.Codename, agp.Description, id)
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
