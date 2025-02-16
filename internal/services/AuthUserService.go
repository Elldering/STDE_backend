package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/hash"
	"STDE_proj/utils/validation"
	"errors"
	"log"
)

//func GetAllAuthUser() ([]models.AuthUser, error) {return repositories.GetAllAuthUser()}

func PostAuthUser(data models.AuthUser) error {
	if !validation.ValidateEmail(data.Email) {
		log.Println("Ошибка: некорректный формат электронной почты")
		return errors.New("некорректный формат электронной почты")
	}
	if !validation.ValidatePhoneNumber(data.PhoneNumber) {
		log.Println("Ошибка: некорректный формат номера телефона")
		return errors.New("некорректный формат номера телефона")
	}

	hashedPassword, err := hash.HashPassword(data.Password)
	if err != nil {
		log.Printf("Ошибка при хешировании пароля: %v", err)
		return err
	}
	data.Password = hashedPassword
	return repositories.PostAuthUser(data)
}

func DeleteAuthUser(id int) error { return repositories.DeleteAuthUser(id) }
