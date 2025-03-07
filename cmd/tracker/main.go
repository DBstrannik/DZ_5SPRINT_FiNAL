package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/actioninfo"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/daysteps"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/trainings"
)

func main() {
	// Создаем объект с персональными данными пользователя.
	person := personaldata.Personal{
		Name:   "Витя",
		Weight: 84.6,
		Height: 1.87,
	}

	// Данные о дневной активности.
	input := []string{
		"678,0h50m",
		"792,1h14m",
		"1078,1h30m",
		"7830,2h40m",
		",3456",
		"12:40:00, 3456",
		"something is wrong",
	}

	fmt.Println("Активность в течение дня")

	// Создаем объект для хранения данных о дневной активности.
	daySteps := daysteps.DaySteps{
		Personal: person,
	}

	// Выводим персональные данные.
	daySteps.Print()

	// Обрабатываем данные о дневной активности.
	actioninfo.Info(input, &daySteps)

	// Данные о тренировках.
	actions := []string{
		"3456,Ходьба,3h00m",
		"something is wrong",
		"678,Бег,0h5m",
		"1078,Бег,0h10m",
		",3456 Ходьба",
		"7892,Ходьба,3h10m",
		"15392,Бег,0h45m",
	}

	// Создаем объект для хранения данных о тренировках.
	trains := trainings.Training{
		Personal: person,
	}

	fmt.Println("Журнал тренировок")

	// Выводим персональные данные.
	trains.Print()

	// Обрабатываем данные о тренировках.
	actioninfo.Info(actions, &trains)
}
