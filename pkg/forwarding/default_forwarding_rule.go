package forwarding

import (
	"strings"

	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/models"
)

type DefaultForwardingRule struct {
	cps []communication.CommunicationProvider
}

func NewForwardingRule(cps []communication.CommunicationProvider) ForwardingRule {
	//todo test this
	if cps == nil {
		cps = []communication.CommunicationProvider{}
	}
	return &DefaultForwardingRule{
		cps: cps,
	}
}

func (r *DefaultForwardingRule) GetProvidersToForwardTo(messageType string) []communication.CommunicationProvider {
	if strings.ToLower(messageType) == models.MESSAGE_TYPE_WARNING {
		cpsToUse := []communication.CommunicationProvider{}
		for i := 0; i < len(r.cps); i++ {
			if strings.ToLower(r.cps[i].GetName()) == models.COMMUNICATIONPROVIDER_DISCORD {
				cpsToUse = append(cpsToUse, r.cps[i])
			}
		}
		return cpsToUse
	} else {
		return []communication.CommunicationProvider{}
	}

}
