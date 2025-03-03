package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
)

func GetAllPositions() ([]models.Position, error) {
	return repositories.GetAllPositions()
}

func GetPositionById(id int) (models.Position, error) {

	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return models.Position{}, errors.New("id не может быть пустой")
	}
	return repositories.GetPositionById(id)
}

func PostPosition(agm models.Position) error {

	err := validation.ValidateEmptyFields(agm.Name, agm.Price)
	if err != nil {
		return errors.New("название и цена не может быть пустой")
	}
	return repositories.PostPosition(agm)
}

func PutPosition(id int, agm models.Position) error {

	err := validation.ValidateEmptyFields(agm.Name, agm.Price, id)
	if err != nil {
		return errors.New("название, цена и id не может быть пустой")
	}

	return repositories.PutPosition(id, agm)
}

func DeletePosition(id int) error {

	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return errors.New("id не может быть пустой")
	}
	return repositories.DeletePosition(id)
}
