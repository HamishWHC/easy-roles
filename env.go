package main

import (
	"encoding/hex"
	"log"
	"os"
)

type DiscordCredentials struct {
	applicationId string
	clientSecret  string
	publicKey     []byte
	botToken      string
}

const DiscordBaseUrl string = "https://discord.com/api"

func GetDiscordCredentials() DiscordCredentials {
	applicationId, exists := os.LookupEnv("APP_ID")
	if !exists {
		log.Fatal("APP_ID must be set.")
	}

	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		log.Fatal("CLIENT_SECRET must be set.")
	}

	publicKeyString, exists := os.LookupEnv("PUBLIC_KEY")
	if !exists {
		log.Fatal("PUBLIC_KEY must be set.")
	}
	publicKey, err := hex.DecodeString(publicKeyString)
	if err != nil || len(publicKeyString) != 64 {
		log.Fatal("PUBLIC_KEY must be a length 64 hex encoded string.")
	}

	botToken, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("BOT_TOKEN must be set.")
	}

	return DiscordCredentials{
		applicationId: applicationId,
		clientSecret:  clientSecret,
		publicKey:     publicKey,
		botToken:      botToken,
	}
}
