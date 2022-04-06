package commands

/*
	コマンドの分類ごとにファイルわけする
*/

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	test  = "test"
	basic = "basic"
)

func init() {
	// コマンドの追加
	commands = append(commands, []*discordgo.ApplicationCommand{
		{
			Name:        test,            // コマンド名
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
		{
			Name:        basic,           // コマンド名
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
	}...)

	// 内容の追加
	commandHandlers[test] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	}
	commandHandlers[basic] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	}
}
