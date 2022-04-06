package commands

/*
	基本的に同一パッケージではprivateとかそういうのないので、
	基本的なものの実装をbaseに書く
*/
import (
	"github.com/bwmarrin/discordgo"
)

var (
	// slashcommandのフロント部分の定義
	commands = make([]*discordgo.ApplicationCommand, 0, 20)

	// slashcommandの内容部分の定義
	commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), 20)
)

// なんか頭大文字だと外部参照可になるらしい
func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

// なんか頭大文字じゃないとだめらしい
func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
