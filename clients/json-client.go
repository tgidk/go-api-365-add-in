package clients

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type JsonClient interface {
	CreateHttpClient(key string)
	GetJson(url string) (interface{}, error)
	GetJsonMap(url string) (map[string]interface{}, error)
}

type jsonClient struct{}

var (
	client *http.Client
	apiKey string
)

func GetNewJsonClient() JsonClient {
	return &jsonClient{}
}

func (*jsonClient) CreateHttpClient(key string) {
	client = &http.Client{Timeout: 10 * time.Second}
	apiKey = key
}

func (*jsonClient) GetJson(url string) (interface{}, error) {
	return nil, nil
}

func (*jsonClient) GetJsonMap(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var target map[string]interface{}
	err = json.Unmarshal(bytes, &target)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return target, err
}
