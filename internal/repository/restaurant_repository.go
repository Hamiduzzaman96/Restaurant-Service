package repository

import (
	"context"

	"github.com/Hamiduzzaman96/Restaurant-Service/internal/domain"
)

type RestaurantRepository interface {
	Create(ctx context.Context, r *domain.Restaurant) error
	GetByID(ctx context.Context, id int64) (*domain.Restaurant, error)
}
