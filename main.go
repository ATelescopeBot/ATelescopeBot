package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ATelescopeBot/ATelescopeBot/events"
	"github.com/ATelescopeBot/ATelescopeBot/handler"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	env := Envvars{}
	env.Init()
	dg, err := discordgo.New("Bot " + env.Token)
	if err != nil {
		log.Fatalf("Can't create Discord session: %s\n", err)
	}

	dg.AddHandler(events.Ready)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(handler.Get)

	dg.Identify.Intents = discordgo.IntentsAll

	if err = dg.Open(); err != nil {
		log.Fatalf("Can't open connection: %s\n", err)
	}

	log.Infoln("========Bot running!========")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer Shutdown(dg)
	defer dg.Close()
}

func Shutdown(s *discordgo.Session) {
	log.Info("Shutdown requested.")
	for _, v := range s.State.Guilds {
		handler.RemoveCommands(s, v)
	}
}
