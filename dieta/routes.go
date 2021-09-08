package dieta

import (
	"fmt"
	"log"

	"github.com/golift/imessage"
)

func (d *Dependencies) SetupRoutes() {
	d.imessage.IncomingCall("(?i)(mangiare)|(?i)dieta", d.DietaRoute)
}

func (d *Dependencies) DietaRoute(msg imessage.Incoming) {
	res := fmt.Sprintf(`%s

Questo era un messaggio automatico.`, d.repo.GetForToday())
	d.imessage.Send(imessage.Outgoing{
		To:   msg.From,
		Text: res,
	})
	d.logger.Printf("Sent message from DietaRoute to %s.", msg.From)
}

type Dependencies struct {
	repo     *dietaRepository
	logger *log.Logger
	imessage *imessage.Messages
}

func NewRoutes(repo *dietaRepository, logger *log.Logger, imessage *imessage.Messages) Dependencies {
	return Dependencies{
		repo:     repo,
		logger: logger,
		imessage: imessage,
	}
}
