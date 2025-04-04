package controllers

import (
    "STDE_proj/internal/services"
    "github.com/gin-gonic/gin"

)

// UpdateMetrics обновляет все метрики, получая данные из слоя сервисов (а там – из репозитория)
func UpdateMetrics(c *gin.Context) {
	// Пользователи
	if count, err := services.GetUserCount(); err == nil {
		services.UserCount.Set(float64(count))
	} 

	// Заказы
	if orderCount, err := services.GetOrderCount(); err == nil {
		services.OrderCount.Set(float64(orderCount))
	} 

	// Заблокированные пользователи
	if blocked, err := services.GetBlockedUserCount(); err == nil {
		services.BlockedUsers.Set(float64(blocked))
	} 

	// Отзывы
	if reviewCount, err := services.GetReviewCount(); err == nil {
		services.ReviewCount.Set(float64(reviewCount))
	} 

	// Средний рейтинг отзывов
	if avgReview, err := services.GetAverageReviewRating(); err == nil {
		services.AverageReviewRating.Set(avgReview)
	} 

	// Средняя цена заказа
	if avgOrderPrice, err := services.GetAverageOrderPrice(); err == nil {
		services.AverageOrderPrice.Set(avgOrderPrice)
	} 

	// Общая выручка
	if totalRevenue, err := services.GetTotalRevenue(); err == nil {
		services.TotalRevenue.Set(totalRevenue)
	} 

	// Количество позиций в заказах
	if orderPositionCount, err := services.GetOrderPositionCount(); err == nil {
		services.OrderPositionCount.Set(float64(orderPositionCount))
	} 

	// Среднее количество товаров в заказе
	if avgOrderItems, err := services.GetAverageOrderItems(); err == nil {
		services.AverageOrderItems.Set(avgOrderItems)
	} 
}
