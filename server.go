package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tgidk/go-api-365-add-in/clients"
	"github.com/tgidk/go-api-365-add-in/controllers"
	router "github.com/tgidk/go-api-365-add-in/http"
	"github.com/tgidk/go-api-365-add-in/mapping"
	"github.com/tgidk/go-api-365-add-in/services"
)

var (
	client          = clients.GetNewJsonClient()
	url             = services.NewUrlService()
	sparkService    = services.NewSparkService(client, url)
	mapper          = mapping.GetNewMapper()
	assetController = controllers.NewAssetController(sparkService, mapper)
	httpRouter      = router.NewMuxRouter()
)

func init() {
	client.CreateHttpClient(os.Getenv("APIKEY"))
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/quote/{id}", assetController.GetAsset)

	err := httpRouter.SERVE(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}
}
