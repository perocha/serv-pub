package controller

import (
	"context"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/gin-gonic/gin"
)

var (
	PUBSUB_NAME  = "orderpubsub"
	PUBSUB_TOPIC = "orders"
)

// Controller example
type Controller struct {
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}

//
func (c *Controller) PostMessage(ctx *gin.Context) {
	log.Printf("PostMessage")

	client, err := dapr.NewClient()
	if err != nil {
		log.Panicf("serv-pub::Fatal error::%s", err)
	}
	defer client.Close()

	context := context.Background()

	myOrder := `{"orderId":"1"}`

	if err := client.PublishEvent(context, PUBSUB_NAME, PUBSUB_TOPIC, &myOrder); err != nil {
		log.Printf("serv-pub::Can't publish event::%s", err)
		ctx.String(http.StatusOK, "Can't publish event")
	} else {
		log.Printf("Published data: %s", myOrder)
		ctx.String(http.StatusOK, "Published data")
	}
}
