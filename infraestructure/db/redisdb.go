package db

import (
	"context"
	"encoding/json"
	"fmt"
	"goml/domain"
	"goml/infraestructure/functions"
	"log"

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

func (dbHandler DBHandler) FindTopProducts() ([]*domain.Product, error) {
	var responseObject []*domain.Product
	//Get
	key := ""
	val, err := dbHandler.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	byteString := []byte(val)
	errUmarshal := json.Unmarshal(byteString, &responseObject)
	if errUmarshal != nil {
		return nil, errUmarshal
	}

	return responseObject, nil
}

func (dbHandler DBHandler) CalculateAndSaveProductsBought(input domain.InputParams) ([]domain.Product, error) {
	var results []domain.Product

	products := functions.MapItems(input)
	results = functions.CalculateByCouponAmount(products, input.CouponAmount)

	//Set
	responseMarshal, _ := json.Marshal(results)
	dbHandler.RedisClient.Set(ctx, "key", string(responseMarshal), 0)
	return results, nil
}

// func (dbHandler DBHandler) FindProductsByClientId(key string) (*domain.Response, error) {
// 	var responseObject domain.Response

// 	//Set
// 	response, err := http.Get("https://api.mercadolibre.com/sites/MCO/search?seller_id=" + key)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}
// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	json.Unmarshal(responseData, &responseObject)
// 	responseMarshal, _ := json.Marshal(responseObject)
// 	dbHandler.RedisClient.Set(ctx, key, string(responseMarshal), 0)

// 	//Get
// 	val, err := dbHandler.RedisClient.Get(ctx, key).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	byteString := []byte(val)
// 	errUmarshal := json.Unmarshal(byteString, &responseObject)
// 	if errUmarshal != nil {
// 		return nil, errUmarshal
// 	}

// 	return &responseObject, nil
// }

// func (dbHandler DBHandler) GetProductsByCouponAndClientId(key string, coupon float64) ([]domain.Product, error) {
// 	var results []domain.Product
// 	var responseObject domain.Response
// 	//Get
// 	val, err := dbHandler.RedisClient.Get(ctx, key).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	byteString := []byte(val)
// 	errUmarshal := json.Unmarshal(byteString, &responseObject)
// 	if errUmarshal != nil {
// 		return nil, errUmarshal
// 	}

// 	results = CalculateByCoupon(responseObject.Results, coupon)
// 	return results, err
// }
