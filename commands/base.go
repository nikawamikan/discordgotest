package commands

/*
	基本的に同一パッケージではprivateとかそういうのないので、
	基本的なものの実装をbaseに書く
*/
import (
	"github.com/bwmarrin/discordgo"
)

type commandBase struct {
	commands        []*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var com commandBase

func init() {
	com.commands = make([]*discordgo.ApplicationCommand, 0, 100)
	com.commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), 100)
}

func (c *commandBase) addCommand(com *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) *commandBase {
	c.commands = append(c.commands, com)
	c.commandHandlers[com.Name] = fn
	return c
}

// なんか頭大文字だと外部参照可になるらしい
func GetCommands() []*discordgo.ApplicationCommand {
	return com.commands
}

// なんか頭大文字じゃないとだめらしい
func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return com.commandHandlers
}
