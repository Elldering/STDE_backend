package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
)

func GetAllUserGroups() ([]models.UserGroup, error) {
	return repositories.GetAllUserGroups()
}

func GetUserGroupById(id int) (models.UserGroup, error) {
	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return models.UserGroup{}, errors.New("id не может быть пустым")
	}
	return repositories.GetUserGroupById(id)
}

func PostUserGroup(agp models.UserGroup) error {
	err := validation.ValidateEmptyFields(agp.Name)
	if err != nil {
		return errors.New("имя не может быть пустым")
	}
	return repositories.PostUserGroup(agp)
}

func PutUserGroup(id int, agp models.UserGroup) error {
	err := validation.ValidateEmptyFields(agp.Name, id)
	if err != nil {
		return errors.New("имя и id не может быть пустым")
	}
	return repositories.PutUserGroup(id, agp)
}

func DeleteUserGroup(id int) error {

	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return errors.New("id не может быть пустым")
	}

	return repositories.DeleteUserGroup(id)
}
