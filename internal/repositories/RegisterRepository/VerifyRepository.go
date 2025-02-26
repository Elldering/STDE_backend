package RegisterRepository

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"errors"
	"log"
)

func CheckAccountVerify(data models.VerifyCode) (bool, error) {

	var isEmailVerified bool = false
	query := "SELECT is_email_verify FROM auth_user WHERE code = $1"
	err := database.DB.QueryRow(query, data.Code).Scan(&isEmailVerified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("почта не подтверждена (аккаунт не зарегистрировал)")
		}
		log.Printf("Ошибка при получении подтверждения почты: %v", err)
		return false, errors.New("ошибка при получении подтверждение почты")
	}
	return isEmailVerified, nil
}

func GetVerificationCode(data models.VerifyCode) (int, error) {

	var authUserID int

	query := "SELECT auth_user_id FROM verify_code WHERE code = $1"
	err := database.DB.QueryRow(query, data.Code).Scan(&authUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("неверный код подтверждения")
		}
		log.Printf("Ошибка при получении кода подтверждения: %v", err)
		return 0, errors.New("ошибка при получении кода подтверждения")
	}
	return authUserID, nil
}

func UpdateEmailVerified(authUserID int) error {
	query := "UPDATE auth_user SET is_email_verify = TRUE WHERE id = $1"
	_, err := database.DB.Exec(query, authUserID)
	if err != nil {
		log.Printf("Ошибка при обновлении статуса подтверждения почты: %v", err)
		return errors.New("ошибка при обновлении статуса подтверждения почты")
	}
	return nil
}

func DeleteVerificationCode(authUserID int) error {
	query := "DELETE FROM verify_code WHERE auth_user_id = $1"
	_, err := database.DB.Exec(query, authUserID)
	if err != nil {
		log.Printf("Ошибка при удалении кода: %v", err)
		return errors.New("ошибка при удалении кода")
	}
	return nil
}
