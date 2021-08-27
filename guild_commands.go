package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	CreateCommandOptions string = `{
		"name": "create",
		"description": "Creates a role menu in the current channel.",
		"options": [
			{
				"type": 1,
				"name": "from-role",
				"description": "Creates a role menu for an existing role.",
				"options": [
					{
						"type": 8,
						"name": "role",
						"description": "The role to create the menu for.",
						"required": true
					}
				]
			},
			{
				"type": 1,
				"name": "new",
				"description": "Creates a role and accompanying role menu for it.",
				"options": [
					{
					"type": 3,
					"name": "role-name",
						"description": "Name for the new role.",
						"required": true
					}
				]
			}
		]
	}`
	GuildApplicationCommandUpsertEndpoint string = DiscordBaseUrl + "/applications/%s/guilds/%s/commands"
)

func UpsertGuildApplicationCommands(credentials DiscordCredentials, guildId string) error {
	client := &http.Client{}

	endpoint := fmt.Sprintf(GuildApplicationCommandUpsertEndpoint, credentials.applicationId, guildId)
	body := strings.NewReader(fmt.Sprintf(
		`[%s]`,
		CreateCommandOptions,
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
