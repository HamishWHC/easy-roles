package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentials := GetDiscordCredentials()

	_, err = discordgo.New(fmt.Sprintf("Bot %s", credentials.botToken))
	if err != nil {
		log.Fatal("Failed to initialise discordgo bot.")
		return
	}

	err = UpsertGlobalApplicationCommands(credentials)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = UpsertGuildApplicationCommands(credentials, "700155253202616392")
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Print("Successfully updated app commands!")

	handler := func(w http.ResponseWriter, r *http.Request) {
		interactionHandler(w, r, credentials)
	}

	http.HandleFunc("/api/interaction", handler)
	log.Print("Serving on http://localhost:8080.")
	err = http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
