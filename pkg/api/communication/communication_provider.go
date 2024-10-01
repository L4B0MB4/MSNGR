package communication

import (
	"context"

	"github.com/L4B0MB4/MSNGR/pkg/models"
)

// interface can be implemented by many providers/channels that want to forward messages
type CommunicationProvider interface {
	GetName() string
	SendMessage(ctx context.Context, messageModel *models.MessageModel) error
}
