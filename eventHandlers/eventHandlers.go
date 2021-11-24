package eventHandlers

import (
	"fmt"

	"github.com/BruceJi7/bummel-bot/config"
	"github.com/BruceJi7/bummel-bot/eventHandlers/events"
	"github.com/bwmarrin/discordgo"
)

func AddEventHandlers(dg *discordgo.Session) {
	fmt.Println("Adding events...")
	fmt.Println("OnReady")
	dg.AddHandler(events.OnReady)

}

func CreateCommands(dg *discordgo.Session) {

	_, err := dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.EraseCommand)
	if err != nil {
		fmt.Println("Error adding erase command:")
		fmt.Println(err)
	} else {
		fmt.Println("Erase command added")
	}
	_, err = dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.ForceLogCommand)
	if err != nil {
		fmt.Println("Error adding forcelog command:")
		fmt.Println(err)
	} else {
		fmt.Println("Forcelog command added")
	}

}
