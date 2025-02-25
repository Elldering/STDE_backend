package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetUserDocuments() ([]models.UserDocument, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, errors.New("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, name, image, is_accepted, type FROM user_document")
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var userDocs []models.UserDocument
	for rows.Next() {
		var ud models.UserDocument
		if err := rows.Scan(&ud.ID, &ud.Name, &ud.Image, &ud.IsAccepted, &ud.Type); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}
		userDocs = append(userDocs, ud)
	}
	return userDocs, nil
}

func GetUserDocumentById(id int) (models.UserDocument, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.UserDocument{}, errors.New("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, name, image, is_accepted, type FROM user_document WHERE id = $1", id)
	var ud models.UserDocument
	if err := row.Scan(&ud.ID, &ud.Name, &ud.Image, &ud.IsAccepted, &ud.Type); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserDocument{}, fmt.Errorf("документ с id %d не найден", id)
		}
		return models.UserDocument{}, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	return ud, nil
}

func PostUserDocument(ud models.UserDocument) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO user_document (name, image, is_accepted, type) VALUES ($1, $2, $3, $4)"
	result, err := db.DB.Exec(query, ud.Name, ud.Image, ud.IsAccepted, ud.Type)
	if err != nil {
		log.Printf("Ошибка при добавлении документа: %v", err)
		return fmt.Errorf("ошибка при добавлении документа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("документ не был добавлен")
	}
	return nil
}

func PutUserDocument(id int, ud models.UserDocument) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "UPDATE user_document SET name = $1, image = $2, is_accepted = $3, type = $4 WHERE id = $5"
	result, err := db.DB.Exec(query, ud.Name, ud.Image, ud.IsAccepted, ud.Type, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении документа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("документ с id %d не найден", id)
	}
	return nil
}

func DeleteUserDocument(id int) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "DELETE FROM user_document WHERE id = $1"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении документа: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("документ с id %d не найден", id)
	}
	return nil
}
