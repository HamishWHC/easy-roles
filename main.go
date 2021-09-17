package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	credentials := GetDiscordCredentials()

	err = UpsertGlobalApplicationCommands(credentials)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = UpsertGuildApplicationCommands(credentials, "700155253202616392")
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Print("Successfully updated app commands!")

	client := NewDiscordClient(credentials)

	handler := func(w http.ResponseWriter, r *http.Request) {
		interactionHandler(w, r, client)
	}

	http.HandleFunc("/api/interaction", handler)
	log.Print("Serving on http://localhost:8080.")
	err = http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
