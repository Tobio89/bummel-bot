package commandHandlers

import (
	"fmt"

	disc "github.com/BruceJi7/bummel-bot/discordHelpers"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers/basicAvailability"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers/erase"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers/scheduleAway"

	"github.com/bwmarrin/discordgo"
)

func AdminCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	data := i.ApplicationCommandData()
	if data.Name != "erase" && data.Name != "force" {
		return
	}
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
	case "force":
		subData := data.Options[0]

		switch subData.Name {
		case "here":
			targetUser := subData.Options[0].UserValue(s)
			basicAvailability.ForceHere(s, i, targetUser)
		case "away":
			targetUser := subData.Options[0].UserValue(s)
			basicAvailability.ForceAway(s, i, targetUser)
		}

	}
}
func ScheduleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	interactionMember := i.Member

	switch data.Name {
	case "here":
		basicAvailability.Here(s, i, interactionMember)

	case "away":
		options := data.Options
		if len(options) == 0 {
			basicAvailability.Away(s, i, interactionMember)
		} else {
			hours := int(options[0].IntValue())
			scheduleAway.AddScheduledAway(hours, &scheduleAway.Schedule)
		}
	}
}
