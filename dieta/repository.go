package dieta

import (
	"io/ioutil"
	"strings"
	"time"
)

type DietaRepository struct {
	dieta []string
}

func NewDietaRepository(path string) DietaRepository {
	daysofweek, err := ioutil.ReadFile(path)
	if err != nil {
		panic(1)
	}

	return DietaRepository{
		dieta: strings.Split(string(daysofweek), "---"),
	}
}

func (d *DietaRepository) GetFromDayOfWeek(daysofweek int) string {
	return strings.TrimSpace(d.dieta[daysofweek-1])
}

func (d *DietaRepository) GetForToday() string {
	return d.GetFromDayOfWeek(int(time.Now().Weekday()))
}
