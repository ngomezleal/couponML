package usecases

import (
	"goml/domain"
	"log"
)

type ProductInteractor struct {
	ProductRepository domain.ProductRepository
}

func NewProductInteractor(repository domain.ProductRepository) ProductInteractor {
	return ProductInteractor{repository}
}

func (pi *ProductInteractor) FindTopProducts() ([]domain.OutputTopProductDto, error) {
	results, err := pi.ProductRepository.FindTopProducts()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return results, err
}

func (pi *ProductInteractor) CalculateAndSaveProductsBought(input domain.InputParams) (domain.OutputProductDto, error) {
	results, err := pi.ProductRepository.CalculateAndSaveProductsBought(input)
	if err != nil {
		log.Println(err.Error())
		return results, err
	}

	return results, err
}
