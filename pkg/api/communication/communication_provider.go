package communication

import (
	"context"

	"github.com/L4B0MB4/MSNGR/pkg/models"
)

type CommunicationProvider interface {
	GetName() string
	SendMessage(ctx context.Context, messageModel *models.MessageModel) error
}
