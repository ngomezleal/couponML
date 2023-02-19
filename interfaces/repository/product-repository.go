package repository

import (
	"goml/domain"
)

type ProductRepository struct {
	handler DBHandler
}

func NewProductRepository(handler DBHandler) ProductRepository {
	return ProductRepository{handler}
}

func (repository ProductRepository) FindProductsByClientId(key string) (*domain.Response, error) {
	results, err := repository.handler.FindProductsByClientId(key)
	if err != nil {
		return nil, err
	}
	return results, err
}

func (repository ProductRepository) GetProductsByCouponAndClientId(key string, coupon float64) ([]domain.Product, error) {
	results, err := repository.handler.GetProductsByCouponAndClientId(key, coupon)
	if err != nil {
		return nil, err
	}
	return results, nil
}
