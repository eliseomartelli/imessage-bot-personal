package ping_pong

import (
	"log"

	"github.com/golift/imessage"
)

func (d *Dependencies) SetupRoutes() {
	d.imessage.IncomingCall("(?i)(/ping)", d.PongRoute)
}

func (d *Dependencies) PongRoute(msg imessage.Incoming) {
	d.imessage.Send(imessage.Outgoing{
		To:   msg.From,
		Text: "Pong!",
	})
	d.logger.Printf("Sent message from PongRoute to %s.", msg.From)
}

type Dependencies struct {
	imessage *imessage.Messages
	logger *log.Logger
}

func NewRoutes(logger *log.Logger, imessage *imessage.Messages) Dependencies {
	return Dependencies{
		logger: logger,
		imessage: imessage,
	}
}
