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

func (pi *ProductInteractor) FindProductsByClientId(key string) (*domain.Response, error) {
	results, err := pi.ProductRepository.FindProductsByClientId(key)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return results, err
}

func (pi *ProductInteractor) GetProductsByCouponAndClientId(key string, coupon float64) ([]domain.Product, error) {
	results, err := pi.ProductRepository.GetProductsByCouponAndClientId(key, coupon)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return results, err
}
