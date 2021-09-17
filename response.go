package main

import (
	"github.com/tidwall/sjson"
)

func MakeInteractionResponse(interactionType int, content string, components string, ephemeral bool) string {
	json, _ := sjson.Set(`{
		"data": {}
	}`, "data.content", content)

	json, _ = sjson.Set(json, "type", interactionType)

	if ephemeral {
		json, _ = sjson.Set(json, "data.flags", 64)
	}

	json, _ = sjson.SetRaw(json, "data.components", components)

	return json
}

func MakeInteractionErrorResponse(msg string) string {
	json, _ := sjson.Set(`{
		"type": 4,
		"data": {"flags": 64}
	}`, "data.content", msg)

	return json
}
