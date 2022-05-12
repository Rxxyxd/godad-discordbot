package handler

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content
	Message := strings.Fields(message)
	if len(Message) < 3 && len(Message) > 1 {
		for i, v := range Message {
			if v == "i'm" || v == "im" {
				noun := "Hi " + Message[i+1] + ", I'm dad"
				_, err := s.ChannelMessageSend(m.ChannelID, noun)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				break
			}

		}
	}

}
