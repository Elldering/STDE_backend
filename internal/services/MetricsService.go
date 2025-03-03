package services

import "STDE_proj/internal/repositories"

type Service struct {
	Repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{Repo: repo}
}

func GetUserCount() (int, error) {
	return repositories.GetUserCount()
}

func GetActiveProjectsCount() (int, error) {
	return repositories.GetActiveProjectsCount()
}

func GetAverageTaskTime() (float64, error) {
	return repositories.GetAverageTaskTime()
}
