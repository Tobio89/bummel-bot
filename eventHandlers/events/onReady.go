package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Activated")
}
