package communication

import "github.com/L4B0MB4/MSNGR/pkg/models"

type CommunicationProvider interface {
	GetName() string
	SendMessage(messageModel *models.MessageModel) error
}
