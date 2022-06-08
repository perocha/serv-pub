package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/perocha/serv-pub/config"
	"github.com/perocha/serv-pub/internal/controller"
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
	log.Printf("Starting %v version: %v port: %v", cfg.App.Name, cfg.App.Version, cfg.App.Port)

	router := gin.Default()
	ctrl := controller.NewController()

	router.Group("/api/v1").POST("/msg", ctrl.PostMessage)
	router.Run(":" + cfg.App.Port)
}
