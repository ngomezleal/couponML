package repository

import (
	"goml/domain"
)

type DBHandler interface {
	FindProductsByClientId(key string) (*domain.Response, error)
	GetProductsByCouponAndClientId(key string, coupon float64) ([]domain.Product, error)
}
