package main

import "fmt"

func makeInteractionResponse(interactionType int, content string, components string, ephemeral bool) string {
	flags := 0
	if ephemeral {
		flags = 64
	}

	return fmt.Sprintf(`{
		"type": %d,
		"data": {
			"content": "%s",
			"components": %s,
			"flags": %d
		}
	}`, interactionType, content, components, flags)
}
