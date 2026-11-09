// Package daysteps предоставляет структуры и методы для парсинга
// и расчета данных о дневных прогулках пользователя
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневных прогулках и личные данные пользователя
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m" и заполняет поля структуры DaySteps
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат строки данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("длительность прогулки должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo формирует и возвращает сводную строку с информацией о дневной прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return info, nil
}
