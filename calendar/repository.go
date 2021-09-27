package calendar

import (
	"os/exec"
)

type calendarRepository struct {
	univesityCalendar string
	workCalendar string
}

func NewRepository(univesityCalendar, workCalendar string) (calendarRepository, error) {
	return calendarRepository{
		univesityCalendar: univesityCalendar,
		workCalendar:      workCalendar,
	}, nil
}

func (d *calendarRepository) GetNowUniversity() (string, bool) {
	return getNow(d.univesityCalendar)
}

func (d *calendarRepository) GetNowWork() (string, bool) {
	return getNow(d.workCalendar)
}

func getNow(calendar string) (string, bool) {
	cmd := exec.Command("icalBuddy", "-ic", calendar, "-b", "- ", "-nc", "eventsNow")
	stdout, err := cmd.Output()
	if err != nil {
		return "", false
	}
	result := string(stdout)
	if len(result) < 1 {
		return result, false
	}
	return result, true
}
