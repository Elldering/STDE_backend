package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
)

func GetReviewsAll() ([]models.Reviews, error) {
	return repositories.GetReviewsAll()
}

func GetReviewsById(id int) (models.Reviews, error) {
	return repositories.GetReviewsById(id)
}

func PostReviews(data models.Reviews) error {
	return repositories.PostReviews(data)
}
