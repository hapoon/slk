package action

import (
	"context"
	"errors"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func ActWebHooks(ctx *cli.Context) (err error) {
	text := ctx.Args().Get(0)
	if text == "" {
		return errors.New("text must be set")
	}

	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		return
	}

	wm := slack.WebhookMessage{
		Username:  cfg.Webhook.Username,
		IconEmoji: cfg.Webhook.IconEmoji,
		IconURL:   cfg.Webhook.IconUrl,
		Channel:   cfg.Webhook.Channel,
		Text:      text,
	}

	if err = slack.PostWebhookContext(context.TODO(), cfg.Webhook.Url, &wm); err != nil {
		return
	}

	return
}
