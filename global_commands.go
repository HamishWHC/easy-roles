package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	RegisterCommandOptions string = `{
		"name": "register",
		"type": 1,
		"description": "Registers the Easy Role commands on your server."
	}`
	GlobalApplicationCommandUpsertEndpoint string = DiscordBaseUrl + "/applications/%s/commands"
)

func UpsertGlobalApplicationCommands(credentials DiscordCredentials) error {
	client := &http.Client{}

	endpoint := fmt.Sprintf(GlobalApplicationCommandUpsertEndpoint, credentials.applicationId)
	body := strings.NewReader(fmt.Sprintf(
		`[%s]`,
		RegisterCommandOptions,
	))

	req, err := http.NewRequest(http.MethodPut, endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", credentials.botToken))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("discord returned non-OK status code: %d, with response body: %s", resp.StatusCode, data)
	}

	return nil
}
