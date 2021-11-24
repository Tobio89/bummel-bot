package commandHandlers

import (
	"fmt"

	"github.com/BruceJi7/bummel-bot/config"
	disc "github.com/BruceJi7/bummel-bot/discordHelpers"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers/erase"

	"github.com/bwmarrin/discordgo"
)

func AdminCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	fmt.Println("Admin command used")
	data := i.ApplicationCommandData()
	options := data.Options

	interactionID := i.Interaction.ID
	interactionChannel, _ := disc.GetChannelByID(s, i.ChannelID)
	interactionMember := i.Member
	interactionMemberCanUseThis, err := disc.MemberHasRoleByName(s, interactionMember, "Botmaster")
	if err != nil {
		fmt.Println("Error on evaluating admin permissions:")
		fmt.Println(err)
	} else {
		if !interactionMemberCanUseThis {
			fmt.Println("You're not allowed")
			return
		}
	}

	switch data.Name {
	case "erase":
		fmt.Println("Erase command used")

		if len(options) == 0 {
			// Triggered single erase mode

			erase.SingleErase(s, i, interactionChannel, interactionID, interactionMember)

		} else {
			// Multiple erase mode:

			erase.MultiErase(s, i, options, interactionChannel, interactionID, interactionMember)

		}

	}
}
func ScheduleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	// interactionID := i.Interaction.ID
	// interactionChannel, _ := disc.GetChannelByID(s, i.ChannelID)
	interactionMember := i.Member

	interactionMemberIsAdmin, err := disc.IsAdmin(s, config.GuildID, interactionMember.User.ID)
	if err != nil {
		fmt.Println("Error on evaluating admin permissions:")
		fmt.Println(err)
	} else {
		if !interactionMemberIsAdmin {
			return
		}
	}

	switch data.Name {
	case "here":
		fmt.Println("Here")

	case "away":
		fmt.Println("Away")
	}
}
