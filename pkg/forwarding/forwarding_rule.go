package forwarding

import "github.com/L4B0MB4/MSNGR/pkg/api/communication"

// used for deciding which communication providers is used when
type ForwardingRule interface {
	GetProvidersToForwardTo(messageType string) []communication.CommunicationProvider
}
