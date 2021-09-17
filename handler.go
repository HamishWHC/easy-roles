package main

import (
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

const (
	InteractionTypePing      = 1
	InteractionTypeCommand   = 2
	InteractionTypeComponent = 3

	InteractionResponseTypePong                        = 1
	InteractionResponseTypeReply                       = 4
	InteractionResponseTypeDelayedReply                = 5
	InteractionResponseTypeComponentMessageEdit        = 7
	InteractionResponseTypeDelayedComponentMessageEdit = 6
)

func interactionHandler(w http.ResponseWriter, r *http.Request, client DiscordClient) {
	if !discordgo.VerifyInteraction(r, ed25519.PublicKey(client.Credentials.publicKey)) {
		fmt.Fprintf(w, "Invalid Discord interaction.")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to respond to interaction: %s", err)
		fmt.Fprintf(w, "Could not read data from request.")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	interaction := gjson.Parse(string(body))

	log.Print(interaction)

	interactionType := interaction.Get("type").Int()

	switch interactionType {
	case InteractionTypePing:
		fmt.Fprintf(w, `{
			"type": 1
		}`)
		return
	case InteractionTypeCommand:
		log.Print("Command recognised!")
		name := interaction.Get("data.name").String()
		handler := NamesToCommandHandlers[name]
		if handler != nil {
			fmt.Fprint(w, handler(interaction, client))
		} else {
			fmt.Fprintf(w, `{
				"type": 4,
				"data": {
					"content": "Something went wrong. Try again later.",
					"flags": 64
				}
			}`)
		}
		return
	case InteractionTypeComponent:
		fmt.Fprintf(w, `{
			"type": 4,
			"data": {
				"content": "Hey thanks for using a component!",
				"flags": 64
			}
		}`)
		return
	default:
		fmt.Fprintf(w, `{
			"type": 4,
			"data": {
				"content": "Something went wrong. Try again later.",
				"flags": 64
			}
		}`)
		return
	}
}
