package calendar

import (
	"fmt"
	"log"
	"sendmessages/hashset"

	"github.com/golift/imessage"
)

const work_route_response string = `In questo momento sto lavorando.

%s

Questo era un messaggio automatico.`

const uni_route_response string = `In questo momento ho questa/e lezione/i

%s

Questo era un messaggio automatico.`

func (d *Dependencies) SetupRoutes() {
	d.imessage.IncomingCall("((?i)(universit.)|(lezione)).*(facendo)?", d.UniveristyRoute)
	d.imessage.IncomingCall("((?i)(lavor.)).*(facendo)?", d.WorkRoute)
}

func (d *Dependencies) WorkRoute(msg imessage.Incoming) {
	d.sendCalendarMessage(msg, work_route_response, d.repo.GetNowWork)
}

func (d *Dependencies) UniveristyRoute(msg imessage.Incoming) {
	d.sendCalendarMessage(msg, uni_route_response, d.repo.GetNowUniversity)
}

func (d *Dependencies) sendCalendarMessage(msg imessage.Incoming, template string, getter func() (string, bool)) {
	if d.allowedContacts.Contains(msg.From) {
		res, present := getter()
		if present {
			res = fmt.Sprintf(template, res)
			d.imessage.Send(imessage.Outgoing{
				To:   msg.From,
				Text: res,
			})
			d.logger.Printf("Sent message from CalendarRoute to %s.", msg.From)
			return
		}
		d.logger.Printf("Message not handled from CalendarRoute to %s.", msg.From)
		return
	}
	d.logger.Printf("Received not allowed request from %s.", msg.From)
}

type Dependencies struct {
	repo            *calendarRepository
	allowedContacts *hashset.Hashset
	logger          *log.Logger
	imessage        *imessage.Messages
}

func NewRoutes(repo *calendarRepository, allowedContacts *hashset.Hashset, logger *log.Logger, imessage *imessage.Messages) Dependencies {
	return Dependencies{
		repo:            repo,
		allowedContacts: allowedContacts,
		logger:          logger,
		imessage:        imessage,
	}
}
