package commands

import "github.com/bwmarrin/discordgo"

var EraseCommand = &discordgo.ApplicationCommand{
	Name:        "erase",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Erase messages in a channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "multiple",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Description: "Specify amount to erase",
		},
	},
}

var HereCommand = &discordgo.ApplicationCommand{
	Name:        "here",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Get the 'here' role",
}
var AwayCommand = &discordgo.ApplicationCommand{
	Name:        "away",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Get the 'away' role",
}
