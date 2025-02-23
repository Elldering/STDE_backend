package RegisterService

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories/RegisterRepository"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"errors"
	"log"
)

func Register(data models.Register) error {
	TypeLogin, err := validation.CheckEmailOrPhoneNumber(data.Login)
	if err != nil {
		return err
	}
	switch TypeLogin {
	case "email":
		if !validation.ValidateEmail(data.Login) {
			return errors.New("некорректная почта")
		}
	case "phone":
		if !validation.ValidatePhoneNumber(data.Login) {
			return errors.New("некорректная номер телефона")
		}
	}
	if !validation.ValidatePassword(data.Password) {
		return errors.New("некорректный пароль")
	}

	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		log.Printf("Ошибка при хешировании пароля: %v", err)
		return err
	}
	data.Password = hashedPassword
	return RegisterRepository.Register(data)
}
