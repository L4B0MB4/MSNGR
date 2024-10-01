package controller

import (
	"github.com/L4B0MB4/MSNGR/pkg/forwarding"
	"github.com/L4B0MB4/MSNGR/pkg/helper"
	"github.com/L4B0MB4/MSNGR/pkg/models"
	"github.com/L4B0MB4/MSNGR/pkg/models/custom_error"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type MessageController struct {
	fwdProvider forwarding.ForwardingProvider
}

func NewMessageController(fwdProvider forwarding.ForwardingProvider) *MessageController {

	return &MessageController{
		fwdProvider: fwdProvider,
	}
}

func (m *MessageController) ForwardMessage(ctx *gin.Context) {
	messageModel := models.MessageModel{}
	err := ctx.ShouldBindBodyWithJSON(&messageModel)
	if err != nil {
		log.Info().Ctx(ctx).Err(err).Msg("Error during model binding")
		helper.AbortWithBadRequest(ctx, err)
		return
	}
	log.Debug().Ctx(ctx).Msg("Received message - forwarding...")
	err = m.fwdProvider.ForwardMessage(ctx, &messageModel)
	if err != nil {

		e1 := err.(*custom_error.NoProvidersError)
		if e1 != nil {
			helper.AbortWithOk(ctx, e1)
			return
		}

		e2 := err.(*custom_error.ForwardFailedError)
		if e2 != nil {
			helper.AbortWithCustomError(ctx, e2)
			return
		}
		helper.AbortWithUnkownError(ctx, err)
		return
	} else {
		ctx.Status(204)
	}
}
