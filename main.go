package main

import (
	"fmt"
	"log"
	"net/http"

	"goml/infraestructure/db"
	"goml/infraestructure/router"
	"goml/interfaces/controllers"
	"goml/interfaces/repository"
	"goml/usecases"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	dbHandler  db.DBHandler
)

func getProductController() controllers.ProductController {
	productRepository := repository.NewProductRepository(dbHandler)
	productInteractor := usecases.NewProductInteractor(productRepository)
	productController := controllers.NewProductController(productInteractor)
	return *productController
}

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})

	var err error
	dbHandler, err = db.NewDBHandler("localhost:6379")
	if err != nil {
		log.Println("Unable to connect to the DataBase")
		return
	}

	productController := getProductController()
	httpRouter.GET("/coupon", productController.Coupon)
	httpRouter.GET("/top", productController.FindTop)
	httpRouter.SERVE(":8000")
}
