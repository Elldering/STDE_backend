package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"STDE_proj/utils/validation"
	"errors"
)

func GetReviewsAll() ([]models.Reviews, error) { return repositories.GetReviewsAll() }

func GetReviewsById(id int) (models.Reviews, error) {
	err := validation.ValidateEmptyFields(id)
	if err != nil {
		return models.Reviews{}, errors.New("id не может быть пустым")
	}
	return repositories.GetReviewsById(id)
}

func PostReviews(data models.Reviews) error {
	err := validation.ValidateEmptyFields(data.AuthUserRecipientID, data.AuthUserSenderID, data.Grade)
	if err != nil {
		return errors.New("id не может быть пустым")
	}
	return repositories.PostReviews(data)
}
