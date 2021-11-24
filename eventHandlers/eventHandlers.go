package eventHandlers

import (
	"fmt"

	"github.com/BruceJi7/bummel-bot/eventHandlers/events"
	"github.com/bwmarrin/discordgo"
)

func AddEventHandlers(dg *discordgo.Session) {
	fmt.Println("Adding events...")
	fmt.Println("OnReady")
	dg.AddHandler(events.OnReady)

}
