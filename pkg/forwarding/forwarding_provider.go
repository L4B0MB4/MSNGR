package forwarding

import (
	"context"

	"github.com/L4B0MB4/MSNGR/pkg/models"
)

type ForwardingProvider interface {
	ForwardMessage(ctx context.Context, messageModel *models.MessageModel) error
}
