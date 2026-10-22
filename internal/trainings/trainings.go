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

// Training хранит данные тренировки и её участника.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку «шаги,тип тренировки,продолжительность».
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("данные тренировки должны содержать три значения")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("продолжительность тренировки: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	t.Steps, t.TrainingType, t.Duration = steps, parts[1], duration
	return nil
}

// ActionInfo возвращает сводку о тренировке.
func (t Training) ActionInfo() (string, error) {
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
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height),
		spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calories), nil
}
