package services

import (
	"STDE_proj/internal/models"
	"STDE_proj/internal/repositories"
	"context"
	"fmt"
)

func GetBasket() ([]models.Basket, error)         { return repositories.GetBasket() }
func GetBasketById(id int) (models.Basket, error) { return repositories.GetBasketById(id) }
func PostBasket(agm models.Basket) error          { return repositories.PostBasket(agm) }
func DeleteBasket(ctx context.Context, id int, isUserID bool) error {
	if isUserID {
		err := repositories.DeleteBasketByUserID(ctx, id)
		if err != nil {
			return fmt.Errorf("ошибка при удалении корзины пользователя с id %d: %w", id, err)
		}
	} else {
		err := repositories.DeleteBasketPosition(ctx, id)
		if err != nil {
			return fmt.Errorf("ошибка при удалении позиции с id %d: %w", id, err)
		}
	}
	return nil
}
