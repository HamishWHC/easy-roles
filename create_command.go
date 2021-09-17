package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const CreateCommandOptions string = `{
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

const CreateGuildRoleEndpoint string = DiscordBaseUrl + "/guilds/%s/roles"

// Should create role menu from either a new role or an existing one.
// If one already exists in the chat, link to it and ask if you want to add to the existing one.

func createHandler(interaction gjson.Result, client DiscordClient) string {
	subCommand := interaction.Get("data.options.0.name").String()

	var role gjson.Result

	switch subCommand {
	case "from-role":
		role = interaction.Get("data.resolved.roles.?")
	case "new":
		guildID := interaction.Get("guild_id").String()
		roleName := interaction.Get("data.options.0.options.0.value").String()
		r, err := client.CreateGuildRole(guildID, roleName, nil, nil, false, true)
		if err != nil {
			log.Print(err)
			return MakeInteractionErrorResponse("An error occurred. Please try again later.")
		}
		role = gjson.Parse(r)
	default:
		log.Print(interaction)
		return MakeInteractionErrorResponse("An error occurred. Please try again later.")
	}

	return MakeInteractionResponse(InteractionResponseTypeReply, "Hey thanks for sending a command!", `[
		{
			"type": 1,
			"components": [
				{
					"custom_id": "button",
					"type": 2,
					"label": "I'm a button!",
					"style": 1
				}
			]
		}
	]`, false)
}
