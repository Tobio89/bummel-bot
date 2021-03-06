package erase

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SingleErase(s *discordgo.Session, i *discordgo.InteractionCreate, interactionChannel *discordgo.Channel, interactionID string, interactionMember *discordgo.Member) {

	err := s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: "Messages Erased", Flags: 1 << 6},
		})

	if err != nil {
		fmt.Println("Error responding to command Erase")
		fmt.Println(err)
	} else {
		fmt.Println("Trigger Erase Command")
		deleteErr := DeleteMessages(1, s, interactionChannel.ID, interactionID)
		if deleteErr != nil {
			fmt.Println("Error deleting one message")
			fmt.Println(deleteErr)
		}
	}

}

func MultiErase(s *discordgo.Session, i *discordgo.InteractionCreate, options []*discordgo.ApplicationCommandInteractionDataOption, interactionChannel *discordgo.Channel, interactionID string, interactionMember *discordgo.Member) {

	eraseAmount := options[0].IntValue()
	fmt.Println(eraseAmount)
	err := s.InteractionRespond(i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: "Messages Erased", Flags: 1 << 6},
		})
	if err != nil {
		fmt.Println("Error responding to command Erase")
		fmt.Println(err)
	} else {
		fmt.Println("Trigger Multiple Erase Command: ", eraseAmount)
		deleteErr := DeleteMessages(int(eraseAmount), s, interactionChannel.ID, interactionID)
		if deleteErr != nil {
			fmt.Println("Error deleting messages")
			fmt.Println(deleteErr)
		}

	}
}

func DeleteMessages(howMany int, s *discordgo.Session, channel string, messageID string) error {

	messages, err := s.ChannelMessages(channel, howMany, messageID, "", "")
	if err != nil {
		fmt.Println("Error getting messages to delete")
		return err
	}
	// fmt.Println(messages)
	var messageIDs []string

	for _, m := range messages {
		messageIDs = append(messageIDs, m.ID)
	}
	messageIDs = append(messageIDs, messageID)

	err = s.ChannelMessagesBulkDelete(channel, messageIDs)
	if err != nil {
		return err
	}

	return nil
}
