package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"utautai/src/bot"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var discord *discordgo.Session

func main() {
	token := os.Getenv("DISCORD_TOKEN")

	fmt.Printf("Token: %s\n", token)

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	comands := bot.NewCommands()

	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("Bot is up!") })
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := comands.HandlerMap[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is up!")
	})

	discord.Identify.Intents = discordgo.IntentsAll

	discord.ApplicationCommandBulkOverwrite(discord.State.SessionID, "", comands.Config)

	err = discord.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
