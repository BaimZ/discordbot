package main

import (
	"log"

	"github.com/baimz/discordbot/bot"
	"github.com/baimz/discordbot/config"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalf("error reading cfg: %s", err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
