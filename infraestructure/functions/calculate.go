package functions

import (
	"goml/domain"
	"math"
	"sort"
)

type ByPrice []domain.Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func CalculateByCouponAmount(items []domain.Product, coupon float64) []domain.Product {
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

	resultS := []domain.Product{}
	var sum = 0.0
	for i := 0; i < len(itemsFilteredByClosest); i++ {
		sum += itemsFilteredByClosest[i].Price
		if sum <= coupon {
			resultS = append(resultS, itemsFilteredByClosest[i])
		} else {
			sumTemp := sum - (math.Abs(itemsFilteredByClosest[i-1].Price + itemsFilteredByClosest[i].Price))
			sumTemp += itemsFilteredByClosest[i].Price
			if sumTemp <= coupon {
				resultS = resultS[:len(resultS)-1]
				resultS = append(resultS, itemsFilteredByClosest[i])
			}
			break
		}
	}
	return resultS
}
