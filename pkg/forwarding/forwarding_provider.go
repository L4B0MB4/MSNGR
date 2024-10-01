package forwarding

import (
	"context"

	"github.com/L4B0MB4/MSNGR/pkg/models"
	"github.com/L4B0MB4/MSNGR/pkg/models/custom_error"
	"github.com/rs/zerolog/log"
)

type ForwardingProvider struct {
	fwdRule ForwardingRule
}

func NewForwardingProvider(fwdRule ForwardingRule) *ForwardingProvider {
	return &ForwardingProvider{
		fwdRule: fwdRule,
	}
}

func (f *ForwardingProvider) ForwardMessage(ctx context.Context, messageModel *models.MessageModel) error {

	cps := f.fwdRule.GetProvidersToForwardTo(messageModel.Type)
	if len(cps) == 0 {
		return custom_error.NewNoProvidersError()
	}
	for _, cp := range cps {
		err := cp.SendMessage(ctx, messageModel)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("Error occured while forwarding message")
			return custom_error.NewForwardFailedError()
		}
	}
	return nil
}
