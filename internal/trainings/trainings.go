package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Training описывает данные о тренировке.
type Training struct {
	Steps        int                   // количество шагов.
	TrainingType string                // тип тренировки (бег или ходьба).
	Duration     time.Duration         // длительность тренировки.
	Personal     personaldata.Personal // персональные данные пользователя.
}

// Print выводит персональные данные пользователя.
func (t Training) Print() {
	fmt.Printf("Имя: %s\nВес: %.1f\nРост: %.2f\n", t.Personal.Name, t.Personal.Weight, t.Personal.Height)
}

// Parse парсит строку с данными о тренировке и заполняет поля структуры Training.
// Формат строки: "количество шагов,тип тренировки,длительность".
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
