package repository

import (
	"goml/domain"
)

type DBHandler interface {
	FindTopProducts() ([]*domain.Product, error)
	CalculateAndSaveProductsBought(input domain.InputParams) ([]domain.Product, error)
}
