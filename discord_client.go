package main

// Weeee!

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/sjson"
)

type DiscordClient struct {
	Credentials DiscordCredentials
	client      http.Client
}

func NewDiscordClient(credentials DiscordCredentials) DiscordClient {
	return DiscordClient{
		Credentials: credentials,
		client:      http.Client{},
	}
}

func (c *DiscordClient) CreateGuildRole(guildID string, name string, permissions *int, color *int, hoist bool, mentionable bool) (string, error) {
	client := http.Client{}

	json, err := sjson.Set("", "name", name)

	if permissions != nil {
		json, err = sjson.Set(json, "permissions", permissions)
		if err != nil {
			return "", err
		}
	}

	if color != nil {
		json, err = sjson.Set(json, "color", color)
		if err != nil {
			return "", err
		}
	}

	json, err = sjson.Set(json, "hoist", hoist)
	if err != nil {
		return "", err
	}

	json, err = sjson.Set(json, "mentionable", mentionable)
	if err != nil {
		return "", err
	}

	log.Print(json)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(CreateGuildRoleEndpoint, guildID), strings.NewReader(json))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", c.Credentials.botToken))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
