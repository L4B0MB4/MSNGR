package forwarding

import (
	"context"

	"github.com/L4B0MB4/MSNGR/pkg/models"
	"github.com/L4B0MB4/MSNGR/pkg/models/custom_error"
	"github.com/rs/zerolog/log"
)

type DefaultForwardingProvider struct {
	fwdRule ForwardingRule
}

func NewForwardingProvider(fwdRule ForwardingRule) *DefaultForwardingProvider {
	return &DefaultForwardingProvider{
		fwdRule: fwdRule,
	}
}

// based on the fowarding rules injected this does commands the communicationProvider to send messages
func (f *DefaultForwardingProvider) ForwardMessage(ctx context.Context, messageModel *models.MessageModel) error {

	cps := f.fwdRule.GetProvidersToForwardTo(messageModel.Type)
	if len(cps) == 0 {
		return custom_error.NewNoProvidersError()
	}
	for _, cp := range cps {
		//to increase speed this could probably go into a go function
		err := cp.SendMessage(ctx, messageModel)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("Error occured while forwarding message")
			return custom_error.NewForwardFailedError()
		}
	}
	return nil
}
