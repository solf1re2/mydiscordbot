package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"strings"
)

// Variables used for command line parameters
var (
	Token   string
	BotID   string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

// launch the discord bot
func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	//dg.AddHandler(messageCreate)
	dg.AddHandler(commandHandler)
	dg.AddHandler(botInit)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// bot setup once the bot is created.
func botInit(s *discordgo.Session, m *discordgo.MessageCreate) {
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// resolve the channel
	// handle private messages - ignore?

	log.Println(m.Content)

	// check for ! at start of line for command input

	// check for @ and username for mention of bot.
	if len(m.Mentions) > 0 && strings.HasPrefix(m.Content, "<@") && m.Mentions[0].ID == s.State.User.ID {
		// message is talking to us. What do we do?

	}

	// echo user
	_, _ = s.ChannelMessageSend(m.ChannelID, m.Content)
	//log.Println(m.Author)
	//log.Println(m.ChannelID)
	//log.Println(m.Message)

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
