package RegisterRepository

import (
	"STDE_proj/utils/database"
	"errors"
	"log"
)

func GetVerificationCode(userID int, code int) (int, error) {
	query := "SELECT code FROM verify_code WHERE auth_user_id = $1"
	err := database.DB.QueryRow(query, userID).Scan(&code)
	if err != nil {
		log.Printf("Ошибка при получении кода подтверждения: %v", err)
		return code, errors.New("ошибка при получении кода подтверждения")
	}
	return code, nil
}

func UpdateEmailVerified(userID int) error {
	query := "UPDATE auth_user SET is_email_verify = TRUE WHERE id = $1"
	_, err := database.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Ошибка при обновлении статуса подтверждения почты: %v", err)
		return errors.New("ошибка при обновлении статуса подтверждения почты")
	}
	return nil
}

func DeleteVerificationCode(userID int) error {
	query := "DELETE FROM verify_code WHERE auth_user_id = $1"
	_, err := database.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Ошибка при удалении кода: %v", err)
		return errors.New("ошибка при удалении кода")
	}
	return nil
}
