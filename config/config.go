package config

import "os"

var Key = os.Getenv("KEY")
var AppID = os.Getenv("APPID")
var GuildID = os.Getenv("GUILDID")
var BotShitChannel = os.Getenv("BOTSHITCHANNEL")
