package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// this is a small tracing implementation that should be enhanced by opentelemetry for proper applications

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId, _ := ctx.Value("spanId").(string)
	e.Str("span-id", spanId)
}

func TracingMiddleWare(c *gin.Context) {
	c.Set("spanId", uuid.New().String())
	c.Next()
}
