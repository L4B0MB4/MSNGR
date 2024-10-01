package api

import (
	"context"
	"net/http"

	"github.com/L4B0MB4/MSNGR/pkg/api/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HttpApi struct {
	httpServer        *http.Server
	router            *gin.Engine
	messageController *controller.MessageController
}

func NewHttpHandler(c *controller.MessageController) *HttpApi {
	//document that gin brings instable go packages with it
	r := gin.Default()
	srv := &http.Server{
		Addr:    "0.0.0.0" + ":" + "6616",
		Handler: r,
	}
	handler := &HttpApi{
		router:            r,
		httpServer:        srv,
		messageController: c,
	}

	handler.registerRoutes()

	return handler
}

func (h *HttpApi) registerRoutes() {
	api := h.router.Group("/api/:tenantId/")
	{

		messageApi := api.Group("messages/")
		{
			messageApi.POST("/", h.messageController.ForwardMessage)
		}
	}

}

func (h *HttpApi) Start() error {
	return h.httpServer.ListenAndServe()
}

func (h *HttpApi) Stop() {
	err := h.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("Error during reading response body")
	}
}
