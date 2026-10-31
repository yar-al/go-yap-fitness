package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yar-al/go-yap-fitness/internal/personaldata"
	"github.com/yar-al/go-yap-fitness/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	var dataPieces = strings.Split(datastring, ",")

	if len(dataPieces) != 2 {
		return errors.New("неверный формат")
	}

	ds.Steps, err = strconv.Atoi(dataPieces[0])
	if err != nil {
		return fmt.Errorf("ошибка конвертации кол-ва шагов: %s", err.Error())
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("некорректное значение кол-ва шагов: %d", ds.Steps)
	}

	ds.Duration, err = time.ParseDuration(dataPieces[1])
	if err != nil {
		return fmt.Errorf("ошибка конвертации длительности: %s", err.Error())
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("некорректное значение длительности: %d", ds.Duration)
	}

	return
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("некорректное кол-во шагов: %d", ds.Steps)
	}
	var stepsLine = fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)

	var distLine = fmt.Sprintf("Дистанция составила %.2f км.\n", spentenergy.Distance(ds.Steps, ds.Height))

	burntCal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка при рассчёте калорий: %s", err)
	}
	var burntCalLine = fmt.Sprintf("Вы сожгли %.2f ккал.\n", burntCal)
	return strings.Join([]string{stepsLine, distLine, burntCalLine}, ""), nil
}
