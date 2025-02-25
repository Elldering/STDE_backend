package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetDocumentAuthUsers() ([]models.DocumentAuthUser, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, errors.New("подключение к базе данных не инициализировано")
	}

	rows, err := db.DB.Query("SELECT id, auth_user_id, user_document_id FROM document_auth_user")
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var docAuthUsers []models.DocumentAuthUser
	for rows.Next() {
		var dau models.DocumentAuthUser
		if err := rows.Scan(&dau.ID, &dau.AuthUserID, &dau.UserDocumentID); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %v", err)
		}
		docAuthUsers = append(docAuthUsers, dau)
	}
	return docAuthUsers, nil
}

func GetDocumentAuthUserById(id int64) (models.DocumentAuthUser, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return models.DocumentAuthUser{}, errors.New("подключение к базе данных не инициализировано")
	}

	row := db.DB.QueryRow("SELECT id, auth_user_id, user_document_id FROM document_auth_user WHERE id = $1", id)
	var dau models.DocumentAuthUser
	if err := row.Scan(&dau.ID, &dau.AuthUserID, &dau.UserDocumentID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.DocumentAuthUser{}, fmt.Errorf("связь документа и пользователя с id %d не найдена", id)
		}
		return models.DocumentAuthUser{}, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	return dau, nil
}

func PostDocumentAuthUser(dau models.DocumentAuthUser) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO document_auth_user (auth_user_id, user_document_id) VALUES ($1, $2)"
	result, err := db.DB.Exec(query, dau.AuthUserID, dau.UserDocumentID)
	if err != nil {
		log.Printf("Ошибка при добавлении связи документа и пользователя: %v", err)
		return fmt.Errorf("ошибка при добавлении связи: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return errors.New("связь документа и пользователя не была добавлена")
	}
	return nil
}

func PutDocumentAuthUser(id int64, dau models.DocumentAuthUser) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "UPDATE document_auth_user SET auth_user_id = $1, user_document_id = $2 WHERE id = $3"
	result, err := db.DB.Exec(query, dau.AuthUserID, dau.UserDocumentID, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении связи документа и пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества измененных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("связь документа и пользователя с id %d не найдена", id)
	}
	return nil
}

func DeleteDocumentAuthUser(id int64) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "DELETE FROM document_auth_user WHERE id = $1"
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении связи документа и пользователя: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("связь документа и пользователя с id %d не найдена", id)
	}
	return nil
}
