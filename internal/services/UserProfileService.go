package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
)

func GetAllUserProfile() ([]models.UserProfile, error) { return repositories.GetAllUserProfile() }

func GetUserProfileById(id int) (models.UserProfile, error) {
	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return models.UserProfile{}, errors.New("id не может быть пустым")
	}
	return repositories.GetUserProfileById(id)
}

func PostUserProfile(agp models.UserProfile) error {
	err := validation.ValidateEmptyFields(agp.FirstName, agp.LastName)
	if err != nil {
		return errors.New("поля Имени и Фамилии являются обязательными для заполнения")
	}
	return repositories.PostUserProfile(agp)
}

func PutUserProfile(id int, agp models.UserProfile) error {
	err := validation.ValidateEmptyFields(agp.FirstName, agp.LastName, id)
	if err != nil {
		return errors.New("поля Имени и Фамилии являются обязательными для заполнения")
	}
	return repositories.PutUserProfile(id, agp)
}

func DeleteUserProfile(id int) error {
	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return errors.New("поле с id не должно быть пустое")
	}
	return repositories.DeleteUserProfile(id)
}
