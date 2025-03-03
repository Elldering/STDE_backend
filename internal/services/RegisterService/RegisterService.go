package RegisterService

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories/RegisterRepository"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"errors"
	"log"
)

func Register(data models.AuthUserRequest) error {
	err := validation.CheckEmailOrPhoneNumber(&data)
	if err != nil {
		return err
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
