package controllers

import (
	"encoding/json"
	"fmt"
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

func (pc *ProductController) Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var params = domain.InputParamsCoupon{}
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	results, err2 := pc.productInteractor.GetProductsByCouponAndClientId(params.Key, params.Coupon)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

func (controller *ProductController) FindAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var params = domain.InputParams{}
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		fmt.Println("error")
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	results, err2 := controller.productInteractor.FindProductsByClientId(params.Key)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}
