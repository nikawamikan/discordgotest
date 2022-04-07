package commands

/*
	コマンドの分類ごとにファイルわけする
*/

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	// コマンドの追加
	com.addCommand(

		// コマンド毎の説明とかはこの辺に書いとけばだいたいわかりそう
		&discordgo.ApplicationCommand{
			Name:        "command",
			Description: "説明",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "test",
				},
			})
		},
	).addCommand( // メソッドチェーン的にすることで多少簡潔に記述可能

		// 2個目
		&discordgo.ApplicationCommand{
			Name:        "command2",
			Description: "説明",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "test",
				},
			})
		},
	)
}
