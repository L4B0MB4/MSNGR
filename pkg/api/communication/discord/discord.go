package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/configuration"
	"github.com/L4B0MB4/MSNGR/pkg/models"
	"github.com/rs/zerolog/log"
)

type DiscordCommunicator struct {
	auth    string
	channel string
}

func NewDiscordCommunicator(config *configuration.ConfigProvider) communication.CommunicationProvider {
	return &DiscordCommunicator{
		auth:    config.GetStringValue("DISCORD_BOT_TOKEN"),
		channel: config.GetStringValue("DISCORD_CHANNEL_ID"),
	}
}

// Returns the name of the provider
func (d *DiscordCommunicator) GetName() string {
	return models.COMMUNICATIONPROVIDER_DISCORD
}

// Sends messages to discord in a template and is part of the interface CommunicationProvider
// Returns unsanitized errors multiple times which need to be handled by the caller
func (d *DiscordCommunicator) SendMessage(ctx context.Context, messageModel *models.MessageModel) error {
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
		log.Info().Ctx(ctx).Err(err).Msg("Could not marshal events")
		return err
	}
	buf := bytes.NewBuffer(bodyBytes)

	httpClient := http.Client{}
	req, err := http.NewRequest("POST", "https://discord.com/api/v10/channels/"+d.channel+"/messages", buf)
	if err != nil {
		log.Warn().Ctx(ctx).Err(err).Msg("Error during request creation")
		return err
	}
	req.Header.Add("Authorization", "Bot "+d.auth)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Warn().Ctx(ctx).Err(err).Msg("Error during response retrival")
		return err
	}
	//for this specific request we know it has to be 200. Otherwise should search in a range of valid status codes
	if resp.StatusCode != 200 {
		log.Warn().Ctx(ctx).Err(err).Msg("Got a non-200 header")
		return fmt.Errorf("unsuccessful request")
	}
	return nil

}
