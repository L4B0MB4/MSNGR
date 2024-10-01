package api

import (
	"context"
	"net/http"

	"github.com/L4B0MB4/MSNGR/pkg/api/controller"
	"github.com/L4B0MB4/MSNGR/pkg/configuration"
	"github.com/L4B0MB4/MSNGR/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HttpApi struct {
	httpServer        *http.Server
	router            *gin.Engine
	messageController *controller.MessageController
}

// Creates a new http api server on 0.0.0.0:6616
// and adds routes for the message endpoint
func NewHttpApi(config *configuration.ConfigProvider, c *controller.MessageController) *HttpApi {
	r := gin.Default()
	srv := &http.Server{
		Addr:    config.GetStringValue("HOST") + ":" + config.GetStringValue("PORT"),
		Handler: r,
	}
	api := &HttpApi{
		router:            r,
		httpServer:        srv,
		messageController: c,
	}

	api.registerRoutes()

	return api
}

func (h *HttpApi) registerRoutes() {
	h.router.Use(helper.TracingMiddleWare)
	api := h.router.Group("/api/:tenantId/")
	{

		messageApi := api.Group("messages/")
		{
			messageApi.POST("/", h.messageController.ForwardMessage)
		}
	}

}

// starts the http api server
func (h *HttpApi) Start() error {
	return h.httpServer.ListenAndServe()
}

// shutsdown the http api server
func (h *HttpApi) Stop() {
	err := h.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("Error during reading response body")
	}
}
