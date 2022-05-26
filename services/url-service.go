package services

import (
	"fmt"

	"github.com/tgidk/go-api-365-add-in/entities"
)

type UrlService interface {
	GetSpark(q entities.SparkQuery) string
}

type urlService struct{}

var (
	//"https://yfapi.net/v8/finance/spark?interval=1d&range=1mo&symbols=INTC"
	sparkUrl = "https://yfapi.net/v8/finance/spark?"
)

func NewUrlService() UrlService {
	return &urlService{}
}

func (*urlService) GetSpark(q entities.SparkQuery) string {
	return fmt.Sprintf("%sinterval=%s&range=%s&symbols=%s", sparkUrl, q.Interval, q.Range, q.Symbol)
}
