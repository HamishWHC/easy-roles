package main

import (
	"github.com/tidwall/gjson"
)

var NamesToCommandHandlers map[string]func(gjson.Result, DiscordClient) string = map[string]func(gjson.Result, DiscordClient) string{
	"create": createHandler,
}
