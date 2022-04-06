package main

import (
	"fmt"
	"os"
	"os/signal"

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

	// コマンドのフロントエンド側の定義
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command", // コマンド名
			Description: "Basic command", // コマンドーの説明
			Options: []*discordgo.ApplicationCommandOption{ // Optionを追加する
				{
					Type:        discordgo.ApplicationCommandOptionString, // オプションの種類
					Name:        "name",                                   // 名前
					Description: "String option",                          // オプションの説明
					Required:    true,                                     // 必須
				},
			},
		},
	}

	// コマンドの中身定義
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			var name string
			msgformat := "test options "

			if option, ok := optionMap["name"]; ok {
				name = option.StringValue()
				msgformat += "\nname: %s"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(msgformat, name),
				},
			})
		},
	}
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
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
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
