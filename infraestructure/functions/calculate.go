package functions

import (
	"goml/domain"
	"math"
	"sort"

	"github.com/peteprogrammer/go-automapper"
)

type ByPrice []domain.Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func CalculateByCouponAmount(items []domain.Product, coupon float64) domain.OutputProductDto {
	items = unique(items)
	sort.Sort(ByPrice(items))

	item := domain.Product{}
	closest := reduce(items, func(acc domain.Product, current domain.Product) domain.Product {
		r := (math.Abs(current.Price-coupon) < math.Abs(acc.Price-coupon))
		if r {
			item = current
			return current
		} else {
			item = acc
			return acc
		}
	}, item)

	itemsFilteredByClosest := filter(items, func(v domain.Product) bool {
		return v.Price <= closest.Price
	})

	products := []domain.Product{}
	var sum = 0.0
	var total = 0.0
	for i := 0; i < len(itemsFilteredByClosest); i++ {
		sum += itemsFilteredByClosest[i].Price
		if sum <= coupon {
			products = append(products, itemsFilteredByClosest[i])
		} else {
			sumTemp := sum - (math.Abs(itemsFilteredByClosest[i-1].Price + itemsFilteredByClosest[i].Price))
			sumTemp += itemsFilteredByClosest[i].Price
			if sumTemp <= coupon {
				products = products[:len(products)-1]
				products = append(products, itemsFilteredByClosest[i])
			}
			break
		}
	}
	for _, item := range products {
		total += item.Price
	}

	productsDto := []domain.ProductDto{}
	automapper.Map(products, &productsDto)
	results := domain.OutputProductDto{
		Items: productsDto,
		Total: total,
	}

	return results
}

func CalculateTopFive(items []*domain.Product) []domain.OutputTopProductDto {
	var occurances = calculateOccurancesOnProducts(items)
	var rows = len(occurances)
	if rows < 5 {
		occurances = occurances[:rows]
	} else {
		occurances = occurances[:5]
	}
	return occurances
}

func calculateOccurancesOnProducts(items []*domain.Product) []domain.OutputTopProductDto {
	var itemsTemp []domain.OutputTopProductDto
	automapper.Map(items, &itemsTemp)

	dict := make(map[domain.OutputTopProductDto]int64)
	for _, item := range itemsTemp {
		dict[item] = dict[item] + 1
	}

	products := make([]domain.OutputTopProductDto, 0, len(dict))
	for item, value := range dict {
		item.Quantity = value
		products = append(products, item)
	}

	sort.Slice(products, func(i, j int) bool {
		return dict[products[i]] > dict[products[j]]
	})

	return products
}
