package info

import (
	"fmt"

	"github.com/ATelescopeBot/ATelescopeBot/handler"
	"github.com/bwmarrin/discordgo"
)

func Ping(ctx handler.Context) {
	ctx.ReplyText("Ping!")
	msg, _ := ctx.Session.InteractionResponse(ctx.Event.Interaction)
	a := msg.Timestamp
	msg, _ = ctx.Session.ChannelMessageEdit(msg.ChannelID, msg.ID, "🏓 Pong!")
	b := msg.EditedTimestamp
	ctx.Session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong! Latency is %vms", b.Sub(a).Milliseconds()))
}

func init() {
	handler.AddCommand(Ping, &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Returns current latency in milliseconds.",
	})
}
