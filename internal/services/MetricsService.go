package services

import "STDE_proj/internal/repositories"

type Service struct {
	Repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetUserCount() (int, error) {
	return s.Repo.GetUserCount()
}
