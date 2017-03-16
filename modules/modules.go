package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/solf1re2/mydiscordbot/modules/plugins"
)

// define interface for modules - commandsList, action struct, init to link to session
type Module interface{}

var (
	pluginCache  map[string]*Plugin
	triggerCache map[string]*Trigger

	PluginList = []Plugin{
		&plugins.RollDice{},
	}
)

type Plugin interface {
	Module
	Commands() []string

	Init(session *discordgo.Session)

	Action(
		command string,
		content string,
		msg *discordgo.Message,
		session *discordgo.Session,
	)
}

type Trigger interface {
	Module
}
