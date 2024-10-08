package helper

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// this is a small tracing implementation that should be enhanced by opentelemetry for proper applications

type TracingHook struct{}

// used for zerolog as a hook whenever "Ctx(ctx)" is used in logging
func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId := GetTraceId(ctx)
	if spanId != "" {
		e.Str("span-id", spanId)

	}
}

func GetTraceId(ctx context.Context) string {
	spanId, _ := ctx.Value("spanId").(string)
	return spanId
}

// Sets a uuid for the context of every gin-request
func TracingMiddleWare(c *gin.Context) {
	c.Set("spanId", uuid.New().String())
	c.Next()
}
