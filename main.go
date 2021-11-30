package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BruceJi7/bummel-bot/config"
	"github.com/BruceJi7/bummel-bot/eventHandlers"
	"github.com/BruceJi7/bummel-bot/eventHandlers/commands/commandHandlers/scheduleAway"
	"github.com/bwmarrin/discordgo"
)

func main() {

	dg, err := discordgo.New("Bot " + config.Key)
	if err != nil {
		fmt.Println("Error starting up:")
		fmt.Println(err)
	}

	//Add all event handlers
	eventHandlers.AddEventHandlers(dg)

	// Initialise slash commands
	eventHandlers.CreateCommands(dg)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsAllWithoutPrivileged
	err = dg.Open() // Open the websocket
	if err != nil {
		fmt.Println("Error initialising websocket:")
		fmt.Println(err)
	}

	go checkSchedule(&scheduleAway.Schedule)

	// Create channel, hold it open
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

func checkSchedule(schedule *[]time.Time) {
	for {
		if len(*schedule) > 0 {
			fmt.Println((*schedule)[0])
		}
	}
}
