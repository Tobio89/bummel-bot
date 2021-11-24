package eventHandlers

import (
	"fmt"

	"github.com/BruceJi7/bummel-bot/config"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers"
	"github.com/BruceJi7/bummel-bot/eventHandlers/events"
	"github.com/bwmarrin/discordgo"
)

func AddEventHandlers(dg *discordgo.Session) {
	fmt.Println("Adding events...")
	fmt.Println("OnReady")
	dg.AddHandler(events.OnReady)

	fmt.Println("Adding command event handlers...")
	fmt.Println("Admin")
	dg.AddHandler(commandHandlers.AdminCommands)
	fmt.Println("Schedule")
	dg.AddHandler(commandHandlers.ScheduleCommands)

}

func CreateCommands(dg *discordgo.Session) {

	_, err := dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.EraseCommand)
	if err != nil {
		fmt.Println("Error adding erase command:")
		fmt.Println(err)
	} else {
		fmt.Println("Erase command added")
	}
	_, err = dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.AdminAvail)
	if err != nil {
		fmt.Println("Error adding admin availability command:")
		fmt.Println(err)
	} else {
		fmt.Println("Availability commands added")
	}
	_, err = dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.HereCommand)
	if err != nil {
		fmt.Println("Error adding here command:")
		fmt.Println(err)
	} else {
		fmt.Println("Here command added")
	}
	_, err = dg.ApplicationCommandCreate(config.AppID, config.GuildID, commands.AwayCommand)
	if err != nil {
		fmt.Println("Error adding away command:")
		fmt.Println(err)
	} else {
		fmt.Println("Away command added")
	}

}
