package action

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
)

type Config struct {
	fp      string         `toml:"-"`
	Webhook WebHooksConfig `toml:"webhook"`
}

type WebHooksConfig struct {
	Url       string `toml:"url"`
	Username  string `toml:"username"`
	IconEmoji string `toml:"icon_emoji"`
	IconUrl   string `toml:"icon_url"`
	Channel   string `toml:"channel"`
}

func (c *Config) Load(profile string) (err error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return
	}
	dir = filepath.Join(dir, "slk")

	var fp string
	if profile == "" {
		fp = filepath.Join(dir, "config.toml")
	} else {
		fp = filepath.Join(dir, fmt.Sprintf("config-%s.toml", profile))
	}
	os.MkdirAll(filepath.Dir(fp), 0700)
	c.fp = fp

	b, err := os.ReadFile(fp)
	if err != nil {
		return
	}

	err = toml.Unmarshal(b, c)

	return
}

func (c Config) Write() (err error) {
	if c.Webhook.Url == "" {
		return errors.New("WebHook URL is empty")
	}

	b, err := toml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(c.fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func ActInit(ctx *cli.Context) (err error) {
	cfg := Config{}
	if err = cfg.Load(ctx.String("profile")); err != nil {
		fmt.Println(err)
		fmt.Println("create new profile")
	}

	s := bufio.NewScanner(os.Stdin)

	fmt.Printf("WebHook URL[%s]:", cfg.Webhook.Url)
	if ok := s.Scan(); !ok {
		log.Fatal(err)
	}
	if s.Text() != "" {
		cfg.Webhook.Url = s.Text()
	}

	fmt.Printf("Username[%s]:", cfg.Webhook.Username)
	if ok := s.Scan(); !ok {
		log.Fatal(err)
	}
	if s.Text() != "" {
		cfg.Webhook.Username = s.Text()
	}

	fmt.Printf("IcomEmoji[%s]:", cfg.Webhook.IconEmoji)
	if ok := s.Scan(); !ok {
		log.Fatal(err)
	}
	if s.Text() != "" {
		cfg.Webhook.IconEmoji = s.Text()
	}

	fmt.Printf("IconURL[%s]:", cfg.Webhook.IconUrl)
	if ok := s.Scan(); !ok {
		log.Fatal(err)
	}
	if s.Text() != "" {
		cfg.Webhook.IconUrl = s.Text()
	}

	fmt.Printf("Channel[%s]:", cfg.Webhook.Channel)
	if ok := s.Scan(); !ok {
		log.Fatal(err)
	}
	if s.Text() != "" {
		cfg.Webhook.Channel = s.Text()
	}

	if err = cfg.Write(); err != nil {
		return
	}

	return
}
