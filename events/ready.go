package events

import (
	"fmt"

	"github.com/ATelescopeBot/ATelescopeBot/handler"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Info("Adding commands...")

	for _, v := range r.Guilds {
		handler.RegisterCommands(s, v)
	}

	go func(s *discordgo.Session, r *discordgo.Ready) {
		for {
			status := fmt.Sprintf("Hey, i'm on %d servers!", len(r.Guilds))
			s.UpdateGameStatus(0, status)
		}
	}(s, r)
}
