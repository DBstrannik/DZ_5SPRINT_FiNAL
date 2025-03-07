package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// DaySteps описывает данные о дневной активности (прогулках).
type DaySteps struct {
	Steps    int                   // количество шагов.
	Duration time.Duration         // длительность прогулки.
	Personal personaldata.Personal // персональные данные пользователя.
}

// Print выводит персональные данные пользователя.
func (ds DaySteps) Print() {
	fmt.Printf("Имя: %s\nВес: %.1f\nРост: %.2f\n", ds.Personal.Name, ds.Personal.Weight, ds.Personal.Height)
}

// Parse парсит строку с данными о прогулке и заполняет поля структуры DaySteps.
// Формат строки: "количество шагов,длительность".
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}

	distance := spentenergy.Distance(ds.Steps)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
