package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tgidk/go-api-365-add-in/entities"
	"github.com/tgidk/go-api-365-add-in/mapping"
	"github.com/tgidk/go-api-365-add-in/services"
)

type assetController struct{}

type AssetController interface {
	GetAsset(response http.ResponseWriter, request *http.Request)
}

var (
	sparkService services.SparkService
	mapper       mapping.Mapper
)

func NewAssetController(s services.SparkService, m mapping.Mapper) AssetController {
	sparkService = s
	mapper = m
	return &assetController{}
}

func (*assetController) GetAsset(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var symbol, exists = mux.Vars(request)["id"]
	if !exists {
		fmt.Println("No symbol found")
		response.WriteHeader((http.StatusNotFound))
		return
	}
	//fmt.Println(symbol)
	//query := entities.SparkQuery{Symbol: "INTC", Interval: "1d", Range: "1mo"}
	query := entities.SparkQuery{Symbol: strings.ToUpper(symbol), Interval: "1d", Range: "1mo"}
	//fmt.Println(query)
	sparks, err := sparkService.GetSpark(query)
	if err != nil {
		//fmt.Println(err.Error())
		response.WriteHeader((http.StatusNotFound))
		return
	}
	asset := mapper.ToAsset(sparks)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(asset)
}
