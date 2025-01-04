package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
)

func Run() {

	discord, err := discordgo.New("Bot " + GetToken())
	HandleError(err)

	discord.ApplicationCommandBulkOverwrite(GetAppId(), GetAppId(), []*discordgo.ApplicationCommand{
		{
			Name:        "call",
			Description: "Вызвать бота в свой голосовой канал.",
		},
	})

	discord.AddHandler(handler)

	err = discord.Open()
	HandleError(err)

	defer func(discord *discordgo.Session) {
		err := discord.Close()
		HandleError(err)
	}(discord)

	fmt.Println("Бот запущен. Нажмите ctrl+C чтобы остановить бота.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func handler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, "asdas")
}
