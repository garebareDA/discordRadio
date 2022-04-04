package main

import (
	"github.com/bwmarrin/discordgo"
)

func Join(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	parsed, err := Parse(m.Content)
	if err != nil {
		return
	}

	if len(parsed) != 1 || parsed[0] != "join" {
		return
	}

	_, err = ChannelVoiceJoin(s, m)
	if err != nil {
		return
	}
}

func ChannelVoiceJoin(s *discordgo.Session, m *discordgo.MessageCreate)(*discordgo.VoiceConnection, error)  {
	for _, g := range s.State.Guilds {
		for _, vs := range g.VoiceStates {
			if m.Author.ID != vs.UserID {
				continue
			}
			return s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, false)
		}
	}

	return nil, nil
}

func ChannelVoiseDisconecct(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	parsed, err := Parse(m.Content)
	if err != nil {
		return
	}

	if len(parsed) != 1 || parsed[0] != "dis" {
		return
	}

	for _, g := range s.State.Guilds {
		for _, vs := range g.VoiceStates {
			if m.Author.ID != vs.UserID {
				continue
			}
			s.VoiceConnections[vs.GuildID].Disconnect()
		}
	}
}
