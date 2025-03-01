package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"fmt"
	"log"
)

// Authenticate аутентифицирует пользователя на основе логина и пароля
func Authentication(data models.AuthUserRequest) error {

	err := validation.CheckEmailOrPhoneNumber(&data)
	if err != nil {
		return err
	}

	user, err := repositories.FindByUsername(data)
	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return fmt.Errorf("неверные учетные данные")
	}
	log.Println(user.ID, user.Login)
	switch data.TypeLogin {
	case "email":
		err = repositories.VerifyEmail(user.Login, user.ID)
		if err != nil {
			return err
		}
	case "phone_number":
		log.Println(data.Login)
	}

	return nil
}
