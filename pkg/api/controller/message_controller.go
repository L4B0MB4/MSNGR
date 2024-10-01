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
	fwdProvider *forwarding.ForwardingProvider
}

func NewMessageController(fwdProvider *forwarding.ForwardingProvider) *MessageController {

	return &MessageController{
		fwdProvider: fwdProvider,
	}
}

func (m *MessageController) ForwardMessage(ctx *gin.Context) {
	messageModel := models.MessageModel{}
	err := ctx.ShouldBindBodyWithJSON(&messageModel)
	if err != nil {
		log.Info().Err(err).Msg("Error during model binding")
		helper.AbortWithBadRequest(ctx, err)
		return
	}
	err = m.fwdProvider.ForwardMessage(&messageModel)
	if err != nil {
		switch e := err.(type) {
		case *custom_error.NoProvidersError:
		case *custom_error.ForwardFailedError:
			helper.AbortWithCustomError(ctx, e)
			return
		default:
			helper.AbortWithUnkownError(ctx, e)
			return
		}
	} else {

	}
}
