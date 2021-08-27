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

func interactionHandler(w http.ResponseWriter, r *http.Request, credentials DiscordCredentials) {
	if !discordgo.VerifyInteraction(r, ed25519.PublicKey(credentials.publicKey)) {
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

	interaction := string(body)

	log.Print(interaction)

	interactionType := gjson.Get(interaction, "type").Int()

	switch interactionType {
	case InteractionTypePing:
		fmt.Fprintf(w, `{
			"type": 1
		}`)
		return
	case InteractionTypeCommand:
		log.Print("Command recognised!")
		response := makeInteractionResponse(InteractionResponseTypeReply, "Hey thanks for sending a command!", `[
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
		fmt.Fprint(w, response)
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
