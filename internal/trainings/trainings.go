// Package trainings предоставляет структуры и методы для парсинга
// и расчета данных о тренировках пользователя
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о конкретной тренировке и личные данные пользователя
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и заполняет поля структуры Training
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("некорректный формат строки данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("длительность тренировки должна быть больше нуля")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает сводную строку с информацией о тренировке
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)

	return info, nil
}
