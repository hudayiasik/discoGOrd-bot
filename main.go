package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string
	config    *config_struct
)

type config_struct struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func read_config() error {
	file, err := ioutil.ReadFile("config.json") // Read the config file
	if err != nil {                             // If there is an error
		return err // Return the error
	}
	if err := json.Unmarshal(file, &config); err != nil {
		return err
	}

	Token = config.Token
	BotPrefix = config.Prefix

	return nil

}

var BotId string
var goBot *discordgo.Session

func start() {
	goBot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID
	goBot.AddHandler(messageCreate)
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == "ping" {
		fmt.Println("Pong!")
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
func main() {
	read_config()
	start()
	<-make(chan struct{})
	return

}
