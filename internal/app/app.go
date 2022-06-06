package app

import (
	"context"
	"log"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/perocha/serv-pub/config"
)

var (
	LOCAL_ENV_FILE = "local.env"
	PUBSUB_NAME    = "orderpubsub"
	PUBSUB_TOPIC   = "orders"
)

//
// Main entry point
//
func Run(cfg *config.Config) {
	log.Printf("Dapr-starter is starting... service name: %v version: %v", cfg.App, cfg.Version)

	serverPort := cfg.App.Port
	log.Printf("serverPort: %s", serverPort)

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	for i := 1; i <= 20; i++ {
		myOrder := `{"orderId":` + strconv.Itoa(i) + `}`
		//		myOrder := order.Order{strconv.Itoa(i), "Description 1", "10,98"}

		// Publish an event using Dapr pub/sub
		if err := client.PublishEvent(ctx, PUBSUB_NAME, PUBSUB_TOPIC, &myOrder); err != nil {
			panic(err)
		}

		log.Printf("Published data: %s", myOrder)

		time.Sleep(2000)
	}

}
