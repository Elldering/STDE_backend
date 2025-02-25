package RegisterService

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories/RegisterRepository"
)

func Verify(data models.VerifyCode) error {
	// Проверяем код подтверждения
	authUserID, err := RegisterRepository.GetVerificationCode(data)
	if err != nil {
		return err
	}

	// Обновляем атрибут подтверждения в таблице auth_user
	err = RegisterRepository.UpdateEmailVerified(authUserID)
	if err != nil {
		return err
	}

	err = RegisterRepository.DeleteVerificationCode(authUserID)
	if err != nil {
		return err
	}
	return nil
}
