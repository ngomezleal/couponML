package repository

import (
	"goml/domain"
)

type DBHandler interface {
	FindTopProducts() ([]domain.OutputTopProductDto, error)
	CalculateAndSaveProductsBought(input domain.InputParams) (domain.OutputProductDto, error)
}
