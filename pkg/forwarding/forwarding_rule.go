package forwarding

import "github.com/L4B0MB4/MSNGR/pkg/api/communication"

type ForwardingRule interface {
	GetProvidersToForwardTo(messageType string) []communication.CommunicationProvider
}
