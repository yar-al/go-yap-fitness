package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yar-al/go-yap-fitness/internal/personaldata"
	"github.com/yar-al/go-yap-fitness/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	var dataPieces = strings.Split(datastring, ",")

	if len(dataPieces) != 3 {
		return errors.New("неверный формат")
	}

	t.Steps, err = strconv.Atoi(dataPieces[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга кол-ва шагов: %s", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("некорректное значение кол-ва шагов: %d", t.Steps)
	}

	t.TrainingType = dataPieces[1]

	t.Duration, err = time.ParseDuration(dataPieces[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %s", err.Error())
	}
	if t.Duration <= 0 {
		return fmt.Errorf("некорректное значение длительности: %d", t.Duration)
	}

	return
}

func (t Training) ActionInfo() (string, error) {
	var typeLine = fmt.Sprintln("Тип тренировки:", t.TrainingType)
	var durLine = fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	var distLine = fmt.Sprintf("Дистанция: %.2f км.\n", spentenergy.Distance(t.Steps, t.Height))
	var speedLine = fmt.Sprintf("Скорость: %.2f км/ч\n", spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration))

	var burntCal float64
	var burntCalErr error

	switch t.TrainingType {
	case "Бег":
		burntCal, burntCalErr = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		burntCal, burntCalErr = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %q", t.TrainingType)
	}
	if burntCalErr != nil {
		return "", fmt.Errorf("ошибка при рассчёте калорий: %s", burntCalErr)
	}

	var burntCalLine = fmt.Sprintf("Сожгли калорий: %.2f\n", burntCal)
	return strings.Join([]string{typeLine, durLine, distLine, speedLine, burntCalLine}, ""), nil
}
