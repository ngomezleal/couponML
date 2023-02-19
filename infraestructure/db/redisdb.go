package db

import (
	"context"
	"encoding/json"
	"fmt"
	"goml/domain"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"

	"github.com/redis/go-redis/v9"
)

type DBHandler struct {
	RedisClient *redis.Client
}

var ctx = context.Background()

func NewDBHandler(connectString string) (DBHandler, error) {
	dbHandler := DBHandler{}
	client := redis.NewClient(&redis.Options{
		Addr:     connectString,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	fmt.Println(pong)
	dbHandler.RedisClient = client
	return dbHandler, nil
}

func (dbHandler DBHandler) FindProductsByClientId(key string) (*domain.Response, error) {
	var responseObject domain.Response

	//Set
	response, err := http.Get("https://api.mercadolibre.com/sites/MCO/search?seller_id=" + key)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &responseObject)
	responseMarshal, _ := json.Marshal(responseObject)
	dbHandler.RedisClient.Set(ctx, key, string(responseMarshal), 0)

	//Get
	val, err := dbHandler.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	byteString := []byte(val)
	errUmarshal := json.Unmarshal(byteString, &responseObject)
	if errUmarshal != nil {
		return nil, errUmarshal
	}

	return &responseObject, nil
}

func (dbHandler DBHandler) GetProductsByCouponAndClientId(key string, coupon float64) ([]domain.Product, error) {
	var results []domain.Product
	var responseObject domain.Response
	//Get
	val, err := dbHandler.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	byteString := []byte(val)
	errUmarshal := json.Unmarshal(byteString, &responseObject)
	if errUmarshal != nil {
		return nil, errUmarshal
	}

	results = CalculateByCoupon(responseObject.Results, coupon)
	return results, err
}

type ByPrice []domain.Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func CalculateByCoupon(items []domain.Product, coupon float64) []domain.Product {
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
