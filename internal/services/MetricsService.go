package services

import "STDE_proj/internal/repositories"

func GetUserCount() (int, error) {
	return repositories.GetUserCount()
}

func GetActiveProjectsCount() (int, error) {
	return repositories.GetActiveProjectsCount()
}

func GetAverageTaskTime() (float64, error) {
	return repositories.GetAverageTaskTime()
}
