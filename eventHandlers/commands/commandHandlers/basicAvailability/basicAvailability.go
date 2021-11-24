package basicAvailability

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/BruceJi7/bummel-bot/config"
	disc "github.com/BruceJi7/bummel-bot/discordHelpers"
)

const hereRoleID string = "900368936716087386"
const awayRoleID string = "900271084149035019"

func replaceRole(s *discordgo.Session, i *discordgo.InteractionCreate, member *discordgo.Member, roleToPut string, roleToTake string, msg string) {
	interactionMemberAlreadyHas := disc.MemberHasRoleByID(member, roleToPut)

	if interactionMemberAlreadyHas {
		err := s.InteractionRespond(i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{Content: "No because that would be double", Flags: 1 << 6},
			})
		if err != nil {
			fmt.Println("Whilst handling replace...")
			fmt.Println("Error responding to command 1")
			fmt.Println(err)
		}
		return
	}
	s.GuildMemberRoleAdd(config.GuildID, member.User.ID, roleToPut)

	//find the away role, remove it from user
	interactionMemberHasRoleToRemove := disc.MemberHasRoleByID(member, roleToTake)
	if interactionMemberHasRoleToRemove {
		s.GuildMemberRoleRemove(config.GuildID, member.User.ID, roleToTake)
	}

	// Respond
	disc.SendMsgToBotChannel(s, (disc.MemberNickOrName(member) + msg))
	err := s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: "Done", Flags: 1 << 6},
		})
	if err != nil {
		fmt.Println("Whilst handling replace command...")
		fmt.Println("Error responding to command 2")
		fmt.Println(err)
	}

}

func Here(s *discordgo.Session, i *discordgo.InteractionCreate, member *discordgo.Member) {
	replaceRole(s, i, member, hereRoleID, awayRoleID, " is up for games")
}

func Away(s *discordgo.Session, i *discordgo.InteractionCreate, member *discordgo.Member) {
	replaceRole(s, i, member, awayRoleID, hereRoleID, " can't play now")
}

func ForceHere(s *discordgo.Session, i *discordgo.InteractionCreate, member *discordgo.User) {

	targetMember, err := disc.FetchMemberByID(s, member.ID)
	if err != nil {
		fmt.Println("Couldn't fetch member")
	}
	replaceRole(s, i, targetMember, hereRoleID, awayRoleID, " was set to here")
}
func ForceAway(s *discordgo.Session, i *discordgo.InteractionCreate, member *discordgo.User) {

	targetMember, err := disc.FetchMemberByID(s, member.ID)
	if err != nil {
		fmt.Println("Couldn't fetch member")
	}
	replaceRole(s, i, targetMember, awayRoleID, hereRoleID, " was set to away")
}
