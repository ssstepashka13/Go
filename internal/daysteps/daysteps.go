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

// DaySteps хранит данные дневной прогулки.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку «шаги,продолжительность».
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("данные прогулки должны содержать два значения")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("продолжительность прогулки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

// ActionInfo возвращает сводку о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, spentenergy.Distance(ds.Steps, ds.Height), calories), nil
}
