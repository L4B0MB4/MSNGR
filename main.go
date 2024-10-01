package main

import (
	"os"

	"github.com/L4B0MB4/MSNGR/pkg/api"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication/discord"
	"github.com/L4B0MB4/MSNGR/pkg/api/controller"
	"github.com/L4B0MB4/MSNGR/pkg/configuration"
	"github.com/L4B0MB4/MSNGR/pkg/forwarding"
	"github.com/L4B0MB4/MSNGR/pkg/helper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger.Debug().Msg("Starting with dependency resolution")
	log.Logger = log.Logger.Hook(helper.TracingHook{})

	config := configuration.NewConfigurationProvider()
	rs := []communication.CommunicationProvider{discord.NewDiscordCommunicator(config)}
	r := forwarding.NewForwardingRule(rs)
	fp := forwarding.NewForwardingProvider(r)
	mCtrl := controller.NewMessageController(fp)
	server := api.NewHttpHandler(config, mCtrl)
	server.Start()
}
