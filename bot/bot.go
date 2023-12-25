package bot

import (
	"fmt"
	"log"

	"github.com/baimz/discordbot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("error new: %s", err.Error())
	}

	u, err := goBot.User("@me")
	if err != nil {
		log.Fatalf("error user: %s", err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatalf("error open: %s", err.Error())
	}

	fmt.Println("bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
