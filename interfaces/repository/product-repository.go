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

func (repository ProductRepository) FindTopProducts() ([]*domain.Product, error) {
	results, err := repository.handler.FindTopProducts()
	if err != nil {
		return nil, err
	}
	return results, err
}

func (repository ProductRepository) CalculateAndSaveProductsBought(input domain.InputParams) ([]domain.Product, error) {
	results, err := repository.handler.CalculateAndSaveProductsBought(input)
	if err != nil {
		return nil, err
	}
	return results, err
}
