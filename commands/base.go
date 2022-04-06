package commnads

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commands = make([]*discordgo.ApplicationCommand, 0, 100)

	// コマンドの中身定義
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), 100)
)

func getCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
