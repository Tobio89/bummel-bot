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

var AdminAvail = &discordgo.ApplicationCommand{
	Name:        "force",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Alter someone else's here or away status",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "here",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Description: "Set someone to be here",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Type:        discordgo.ApplicationCommandOptionUser,
					Description: "Who to modify",
					Required:    true,
				},
			},
		},
		{
			Name:        "away",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Description: "Set someone to be away",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Type:        discordgo.ApplicationCommandOptionUser,
					Description: "Who to modify",
					Required:    true,
				},
			},
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
