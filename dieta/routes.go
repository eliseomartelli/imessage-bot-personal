package dieta

import (
	"fmt"

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
}

type Dependencies struct {
	repo     *DietaRepository
	imessage *imessage.Messages
}

func New(repo *DietaRepository, imessage *imessage.Messages) Dependencies {
	return Dependencies{
		repo:     repo,
		imessage: imessage,
	}
}
