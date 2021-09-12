package dieta

import (
	"io/ioutil"
	"strings"
	"time"
)

type dietaRepository struct {
	dieta []string
}

func NewRepository(path string) (dietaRepository, error) {
	daysofweek, err := ioutil.ReadFile(path)
	if err != nil {
		return dietaRepository{}, err
	}

	return dietaRepository{
		dieta: strings.Split(string(daysofweek), "---"),
	}, nil
}

func (d *dietaRepository) getFromDayOfWeek(daysofweek int) string {
	return strings.TrimSpace(d.dieta[daysofweek-1])
}

func (d *dietaRepository) GetForToday() string {
	return d.getFromDayOfWeek(int(time.Now().Weekday()))
}
