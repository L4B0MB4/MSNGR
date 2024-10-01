package main

import (
	"os"

	"github.com/L4B0MB4/MSNGR/pkg/api"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication/discord"
	"github.com/L4B0MB4/MSNGR/pkg/api/controller"
	"github.com/L4B0MB4/MSNGR/pkg/forwarding"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger.Debug().Msg("Starting with dependency resolution")

	rs := []communication.CommunicationProvider{discord.NewDiscordCommunicator()}
	r := forwarding.NewForwardingRule(rs)
	fp := forwarding.NewForwardingProvider(r)
	mCtrl := controller.NewMessageController(fp)
	server := api.NewHttpHandler(mCtrl)
	server.Start()
}
