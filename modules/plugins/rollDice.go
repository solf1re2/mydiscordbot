package plugins

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type RollDice struct{}

func (r *RollDice) Commands() []string {
	return []string{
		"roll",
	}
}

func (r *RollDice) Init(session *discordgo.Session) {}

func (r *RollDice) Action(command string, content string, msg *discordgo.Message, session *discordgo.Session) {
	defer func() {
		err := recover()
		if err != nil {
			session.ChannelMessageSend(msg.ChannelID, "I've lost my dice!")
		}
	}()
	d := 2
	session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Roll : %v", d))
}
