package discordHelpers

import (
	"errors"
	"fmt"

	"github.com/BruceJi7/bummel-bot/config"

	"github.com/bwmarrin/discordgo"
)

type LogPrefixes struct {
	Error       string
	Forcelog    string
	EraseOne    string
	EraseMulti  string
	RoleAdded   string
	RoleRemoved string
}

// // Get channel by name
func GetChannelByName(s *discordgo.Session, name string) (c *discordgo.Channel, err error) {
	channels, _ := s.GuildChannels(config.GuildID)
	for _, c := range channels {
		if c.Name == name {
			return c, nil
		}
	}
	return channels[0], errors.New("channel not found")
}

func GetChannelByID(s *discordgo.Session, cID string) (c *discordgo.Channel, err error) {

	channels, _ := s.GuildChannels(config.GuildID)
	for _, c := range channels {
		if c.ID == cID {
			return c, nil
		}
	}
	return channels[0], errors.New("channel not found")
}

func GetRoleByName(s *discordgo.Session, roleName string) (role *discordgo.Role, err error) {
	roles, _ := s.GuildRoles(config.GuildID)

	for _, role := range roles {
		if role.Name == roleName {
			return role, nil
		}
	}
	return roles[0], errors.New("role not found")
}
func GetRoleByID(s *discordgo.Session, roleID string) (role *discordgo.Role, err error) {
	roles, _ := s.GuildRoles(config.GuildID)

	for _, role := range roles {
		if role.ID == roleID {
			return role, nil
		}
	}
	return roles[0], errors.New("role not found")
}

func MemberHasRoleByName(s *discordgo.Session, mem *discordgo.Member, roleName string) (bool, error) {
	roleObj, err := GetRoleByName(s, roleName)
	if err != nil {
		return false, err
	}

	for _, role := range mem.Roles {
		if role == roleObj.ID {
			return true, nil
		}
	}
	return false, nil
}

func FetchMember(s *discordgo.Session, userDetails string) (member *discordgo.Member, err error) {
	guildMembers, err := s.GuildMembers(config.GuildID, "", 1000)

	if err != nil {
		fmt.Println("Error finding member")
		fmt.Println(err)
	}
	for _, member := range guildMembers {
		if member.User.ID == userDetails {
			return member, nil
		}
	}
	return guildMembers[0], errors.New("member not found")
}

func IsAdmin(s *discordgo.Session, guildID string, userID string) (bool, error) {
	return memberHasPermission(s, guildID, userID, discordgo.PermissionAdministrator)
}

func memberHasPermission(s *discordgo.Session, guildID string, userID string, permission int64) (bool, error) {
	member, err := s.State.Member(guildID, userID)
	if err != nil {
		if member, err = s.GuildMember(guildID, userID); err != nil {
			return false, err
		}
	}

	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range member.Roles {
		role, err := s.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}

	return false, nil
}
