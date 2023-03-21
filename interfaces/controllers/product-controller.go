package controllers

import (
	"encoding/json"
	"goml/domain"
	"goml/usecases"
	"net/http"
)

type ProductController struct {
	productInteractor usecases.ProductInteractor
}

func NewProductController(productInteractor usecases.ProductInteractor) *ProductController {
	return &ProductController{productInteractor}
}

func (pc *ProductController) Coupon(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var params = domain.InputParams{}
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	results, err2 := pc.productInteractor.CalculateAndSaveProductsBought(params)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

func (controller *ProductController) FindTop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	results, err2 := controller.productInteractor.FindTopProducts()
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
