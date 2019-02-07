package Twitch

import (
	"github.com/genesixx/coalibot/Struct"

	"github.com/nlopes/slack"
)

func Emotes(option string, event *Struct.Message) bool {
	params := slack.FileUploadParameters{
		File:     "emotes/" + option + ".png",
		Channels: []string{event.Channel},
	}
	event.API.UploadFile(params)
	return true
}
