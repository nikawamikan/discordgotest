package commands

/*
	基本的に同一パッケージではprivateとかそういうのないので、
	基本的なものの実装をbaseに書く
*/
import (
	"github.com/bwmarrin/discordgo"
)

// structの方が微妙に使いやすいので定義（java民の感想）
type commandBase struct {
	commands        []*discordgo.ApplicationCommand
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// コマンドのフロント側の名称をそのまま関数のMapに使用してsetします
func (c *commandBase) addCommand(com *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) *commandBase {
	c.commands = append(c.commands, com)
	c.commandHandlers[com.Name] = fn
	return c
}

// 空のオブジェクト定義
var com commandBase

// 空の配列長無駄に100定義してるけど全然たぶん必要ないので考えとく（たぶんそこまでパフォーマンスに影響でないので0でも良いとは思う）
func init() {
	com.commands = make([]*discordgo.ApplicationCommand, 0, 100)
	com.commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), 100)
}

// なんか頭大文字だと外部参照可になるらしい
func GetCommands() []*discordgo.ApplicationCommand {
	return com.commands
}

// なんか頭大文字じゃないとだめらしい
func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return com.commandHandlers
}
