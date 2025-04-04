package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/hash"
	"fmt"
	"log"
)

// Authenticate аутентифицирует пользователя на основе логина и пароля
func Authentication(data models.AuthUserRequest) error {

	

	user, err := repositories.FindByUsername(data)
	if err != nil {
		return fmt.Errorf("ошибка при поиске пользователя: %v", err)
	}
	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return fmt.Errorf("неверные учетные данные")
	}
	log.Println(user.ID, user.Login)
	repositories.VerifyEmail(user.Login, user.ID)

	return nil
}
