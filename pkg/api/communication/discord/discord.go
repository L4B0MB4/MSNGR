package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/models"
	"github.com/rs/zerolog/log"
)

type DiscordCommunicator struct {
}

func NewDiscordCommunicator() communication.CommunicationProvider {
	return &DiscordCommunicator{}
}

// GetName implements CommunicationProvider.
func (d *DiscordCommunicator) GetName() string {
	return models.COMMUNICATIONPROVIDER_DISCORD
}

// SendMessage implements CommunicationProvider.
func (d *DiscordCommunicator) SendMessage(messageModel *models.MessageModel) error {
	msgEmbed := messageEmbed{
		Title:       messageModel.Name,
		Description: messageModel.Description,
	}
	msgTemplate := messageTemplate{
		Content: "Forwarded message",
		Embeds:  []messageEmbed{msgEmbed},
	}
	bodyBytes, err := json.Marshal(msgTemplate)
	if err != nil {
		log.Info().Err(err).Msg("Could not marshal events")
		return err
	}
	buf := bytes.NewBuffer(bodyBytes)

	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://discord.com/api/v10/channels/.../messages", buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	//for this specific request we know it has to be 200. Otherwise should search in a range of valid status codes
	if resp.StatusCode != 200 {
		log.Info().Err(err).Msg("Got non 2XX header")
		return fmt.Errorf("UNSUCCESSFUL REQUEST")
	}
	return nil

}
