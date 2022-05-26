package services

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/tgidk/go-api-365-add-in/clients"
	"github.com/tgidk/go-api-365-add-in/entities"
)

type SparkService interface {
	GetSpark(q entities.SparkQuery) (entities.Spark, error)
}

type sparkService struct{}

var (
	client clients.JsonClient
	url    UrlService
)

func NewSparkService(c clients.JsonClient, u UrlService) SparkService {
	client = c
	url = u
	return &sparkService{}
}

func (*sparkService) GetSpark(q entities.SparkQuery) (entities.Spark, error) {
	_url := url.GetSpark(q)
	spark := entities.Spark{}
	fmt.Println(spark)

	target, err := client.GetJsonMap(_url)
	if err != nil {
		log.Fatal(err)
		return spark, err
	}
	//fmt.Print(target)
	var data = target[strings.ToUpper(q.Symbol)]
	if data == nil {
		//map[spark:map[error:map[code:Not Found description:No data found for spark symbols] result:<nil>]]
		return spark, errors.New("data not found for symbol")
	}
	quote := data.(map[string]interface{})
	err = mapstructure.Decode(quote, &spark)
	if err != nil {
		log.Fatal(err)
		return spark, err
	}

	return spark, err
}

// fmt.Println("l spark = ", reflect.TypeOf(spark))
// fmt.Println(spark)
// fmt.Println("print key value pairs:")
// for key, value := range quote {
// 	// Each value is an interface{} type, that is type asserted as a string
// 	fmt.Println(key, value)
// 	fmt.Println("key = ", reflect.TypeOf(key))
// 	fmt.Println("value = ", reflect.TypeOf(value))
// 	fmt.Println("------------------------")
// }
