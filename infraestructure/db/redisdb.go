package db

import (
	"context"
	"encoding/json"
	"fmt"
	"goml/domain"
	"goml/infraestructure/functions"
	"log"

	"github.com/google/uuid"
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

func (dbHandler DBHandler) FindTopProducts() ([]domain.OutputTopProductDto, error) {
	var responseObject []domain.OutputTopProductDto
	var temporalProducts []*domain.Product
	var responseUnMarshalObject []*domain.Product

	//Get
	keys := getAllKeys(dbHandler, ctx, "client*")
	for _, key := range keys {
		val, err := dbHandler.RedisClient.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		byteString := []byte(val)
		errUmarshal := json.Unmarshal(byteString, &responseUnMarshalObject)
		if errUmarshal != nil {
			return nil, errUmarshal
		}

		for _, item := range responseUnMarshalObject {
			temporalProducts = append(temporalProducts, item)
		}
	}

	//Get Top Five
	responseObject = functions.CalculateTopFive(temporalProducts)
	return responseObject, nil
}

func (dbHandler DBHandler) CalculateAndSaveProductsBought(input domain.InputParams) (domain.OutputProductDto, error) {
	var results domain.OutputProductDto

	products := functions.MapItems(input)
	results = functions.CalculateByCouponAmount(products, input.CouponAmount)

	//Set
	id := uuid.New()
	responseMarshal, _ := json.Marshal(results.Items)
	dbHandler.RedisClient.Set(ctx, "client-"+id.String(), string(responseMarshal), 0)
	return results, nil
}

func getAllKeys(dbHandler DBHandler, ctx context.Context, key string) []string {
	keys := []string{}

	iter := dbHandler.RedisClient.Scan(ctx, 0, key, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	return keys
}
