package RegisterService

import (
	"STDE_proj/internal/repositories/RegisterRepository"
	"errors"
)

func Verify(id int, code int) error {
	// Проверяем код подтверждения
	storedCode, err := RegisterRepository.GetVerificationCode(id, code)
	if err != nil {
		return err
	}
	if storedCode != code {
		return errors.New("неверный код подтверждения")
	}

	// Обновляем атрибут подтверждения в таблице auth_user
	err = RegisterRepository.UpdateEmailVerified(id)
	if err != nil {
		return err
	}

	err = RegisterRepository.DeleteVerificationCode(id)
	if err != nil {
		return err
	}
	return nil
}
