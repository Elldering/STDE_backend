package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func VerifyService(data models.VerifyCode) error {
	// Проверяем код подтверждения
	authUserID, err := repositories.GetVerificationCode(data)
	if err != nil {
		return err
	}

	//isEmailVerify, err := RegisterRepository.CheckAccountVerify(data)
	//if err != nil {
	//	return err
	//}

	if data.Type == "reg" {
		err = repositories.UpdateEmailVerified(authUserID)
		if err != nil {
			return err
		}
	}

	err = repositories.DeleteVerificationCode(authUserID)
	if err != nil {
		return err
	}
	return nil
}
