package dieta

import (
	"io/ioutil"
	"strings"
	"time"
)

type dietaRepository struct {
	dieta []string
}

func NewRepository(path string) dietaRepository {
	daysofweek, err := ioutil.ReadFile(path)
	if err != nil {
		panic(1)
	}

	return dietaRepository{
		dieta: strings.Split(string(daysofweek), "---"),
	}
}

func (d *dietaRepository) getFromDayOfWeek(daysofweek int) string {
	return strings.TrimSpace(d.dieta[daysofweek-1])
}

func (d *dietaRepository) GetForToday() string {
	return d.getFromDayOfWeek(int(time.Now().Weekday()))
}
