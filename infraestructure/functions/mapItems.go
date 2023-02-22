package functions

import (
	"encoding/json"
	"fmt"
	"goml/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func MapItems(input domain.InputParams) []domain.Product {
	var responseObject domain.Product
	var products []domain.Product

	for _, v := range input.Items {
		response, err := http.Get("https://api.mercadolibre.com/items/" + v)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		json.Unmarshal(responseData, &responseObject)
		products = append(products, responseObject)

		if err != nil {
			log.Fatal(err)
		}
	}
	return products
}
