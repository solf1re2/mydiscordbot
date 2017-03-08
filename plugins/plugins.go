package plugins

// define interface for plugins - commandsList, action struct, init to link to session
type Plugin interface {
	Commands() []string
}
