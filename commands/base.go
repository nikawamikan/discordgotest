package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// slashcommandのフロント部分の定義
	commands = make([]*discordgo.ApplicationCommand, 0, 20)

	// slashcommandの内容部分の定義
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), 20)
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
