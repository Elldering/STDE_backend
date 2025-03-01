package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
)

func GetPermissions() ([]models.Permission, error) {
	return repositories.GetAllPermissions()
}

func GetPermissionById(id int) (models.Permission, error) {
	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return models.Permission{}, errors.New("id не может быть пустой")
	}
	return repositories.GetPermissionById(id)
}

func PostPermission(agp models.Permission) error {

	err := validation.ValidateEmptyFields(agp.Codename, agp.Description)
	if err != nil {
		return errors.New("название и описание не может быть пустой")
	}
	return repositories.PostPermission(agp)
}

func PutPermission(id int, agp models.Permission) error {

	err := validation.ValidateEmptyFields(agp.Codename, agp.Description, id)
	if err != nil {
		return errors.New("название, id и описание не может быть пустой")
	}
	return repositories.PutPermission(id, agp)
}

func DeletePermission(id int) error {

	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return errors.New("id не может быть пустой")
	}
	return repositories.DeletePermission(id)
}
