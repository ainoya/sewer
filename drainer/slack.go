package drainer

import (
	"fmt"
	"os"

	"github.com/bluele/slack"
)

type SlackDrainer struct {
	channel   string
	webHook   *slack.WebHook
	iconEmoji string
}

func NewSlackDrainer() (*SlackDrainer, error) {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")

	if webhookURL == "" {
		return nil, fmt.Errorf("environment variable SLACK_WEBHOOK_URL is not found")
	}

	hook := slack.NewWebHook(webhookURL)

	channel := os.Getenv("SLACK_CHANNEL")
	if channel == "" {
		return nil, fmt.Errorf("environment variable SLACK_CHANNEL is not found")
	}

	iconEmoji := os.Getenv("SLACK_ICON_EMOJI")

	return &SlackDrainer{
		channel:   channel,
		webHook:   hook,
		iconEmoji: iconEmoji,
	}, nil
}

func (d SlackDrainer) Drain(message string) error {
	pl := &slack.WebHookPostPayload{
		Channel: d.channel,
		Text:    message,
	}
	if d.iconEmoji != "" {
		pl.IconEmoji = d.iconEmoji
	}

	return d.webHook.PostMessage(pl)
}
