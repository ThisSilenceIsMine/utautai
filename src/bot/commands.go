package bot

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {

	userId := i.Member.User.ID

	vs, err := s.State.VoiceState(i.GuildID, userId)

	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You are not in a voice channel!",
			},
		})
		return
	}

	s.ChannelVoiceJoin(i.GuildID, vs.ChannelID, false, false)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}

func Amongus(s *discordgo.Session, i *discordgo.InteractionCreate) {

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Amongus!",
		},
	})
}

type Commands struct {
	Config     []*discordgo.ApplicationCommand
	HandlerMap map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommands() *Commands {
	return &Commands{
		Config:     GetCommands(),
		HandlerMap: GetCommandHandlers(),
	}
}

func GetCommands() []*discordgo.ApplicationCommand {

	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping!",
			Type:        discordgo.ChatApplicationCommand,
		},
		{
			Name:        "amongus",
			Description: "Amongus!",
			Type:        discordgo.ChatApplicationCommand,
		},
	}
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping":    Ping,
		"amongus": Amongus,
	}
}
