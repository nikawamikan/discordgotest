package main

/*
	mainにイベントの呼び出しとかを定義
*/
import (
	"fmt"
	"os"
	"os/signal"
	"testbot/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// .envファイルで定義します
var (
	GuildID        string
	BotToken       string
	RemoveCommands bool
)

var s *discordgo.Session

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Dotenv read error: %v", err)
	}

	GuildID = os.Getenv("GUILDID")
	BotToken = os.Getenv("TOKEN")
	RemoveCommands = os.Getenv("REMOVECOMMAND") == "true"

	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Printf("Invalid bot parameters: %v", err)
	}
}

var (
	commandList     = commands.GetCommands()
	commandHandlers = commands.GetCommandHandlers()
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("fmtged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		fmt.Printf("Cannot open the session: %v", err)
	}

	fmt.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commandList))
	for i, v := range commandList {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			fmt.Printf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Press Ctrl+C to exit")
	<-stop

	if RemoveCommands {
		fmt.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
			if err != nil {
				fmt.Printf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	fmt.Println("Gracefully shutting down.")
}
