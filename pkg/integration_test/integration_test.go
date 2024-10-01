package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/L4B0MB4/MSNGR/pkg/api"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication"
	"github.com/L4B0MB4/MSNGR/pkg/api/communication/discord"
	"github.com/L4B0MB4/MSNGR/pkg/api/controller"
	"github.com/L4B0MB4/MSNGR/pkg/configuration"
	"github.com/L4B0MB4/MSNGR/pkg/forwarding"
	"github.com/L4B0MB4/MSNGR/pkg/models/custom_error"
)

func Setup() *api.HttpApi {
	config := configuration.NewConfigurationProvider()
	rs := []communication.CommunicationProvider{discord.NewDiscordCommunicator(config)}
	r := forwarding.NewForwardingRule(rs)
	fp := forwarding.NewForwardingProvider(r)
	mCtrl := controller.NewMessageController(fp)
	server := api.NewHttpHandler(config, mCtrl)
	return server
}

// These tests are not fully integration due to their lack of checking if the discord message list did change
// but they test for a successful discord client call
// also: make sure to have env-file/env-variables set
func TestIfWarningMessageGetsForwarded(t *testing.T) {
	s := Setup()
	go func() {
		s.Start()
	}()

	payload := strings.NewReader(`{ 
		"Type": "Warning", 
		"Name": "Backup Failure", 
		"Description": "The backup failed due to a database problem"
	} `)

	client := &http.Client{}
	url := "http://localhost:6616/api/1/messages"
	method := "POST"
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	if res.StatusCode != 204 {
		t.Error(err)
		t.FailNow()
		return
	}

	s.Stop()
}
func TestIfInfoMessageDoesNotGetForwarded(t *testing.T) {

	s := Setup()
	go func() {
		s.Start()
	}()

	payload := strings.NewReader(`{ 
		"Type": "Warning!", 
		"Name": "Backup Failure", 
		"Description": "The backup failed due to a database problem"
	} `)

	client := &http.Client{}
	url := "http://localhost:6616/api/1/messages"
	method := "POST"
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if res.StatusCode != 200 {
		t.Error(err)
		t.FailNow()
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if !strings.Contains(string(body), custom_error.NewNoProvidersError().Error()) {

		t.Error(fmt.Errorf("info message about missing providers is missing"))
		t.FailNow()
	}

	s.Stop()

}

func TestIfBadRequestGetsRejected(t *testing.T) {

	s := Setup()
	go func() {
		s.Start()
	}()

	payload := strings.NewReader(`{ 
		"Type": "Warning!", 
		"Name": 123, 
		"Description": "The backup failed due to a database problem"
	} `)

	client := &http.Client{}
	url := "http://localhost:6616/api/1/messages"
	method := "POST"
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return
	}

	if res.StatusCode != 400 {
		t.Error(err)
		t.FailNow()
		return
	}

	s.Stop()
}
