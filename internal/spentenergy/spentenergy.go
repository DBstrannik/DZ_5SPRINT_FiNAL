package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// Distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
// Для расчета дистанции нужно шаги умножить на длину шага lenStep и разделить на mInKm.
func Distance(steps int) float64 {
	return float64(steps) * lenStep / mInKm
}

// MeanSpeed возвращает значение средней скорости движения во время тренировки.
// Параметры:
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps)
	return distance / duration.Hours()
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
// Параметры:
// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("weight and height must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}
	speed := MeanSpeed(steps, duration)
	return (walkingCaloriesWeightMultiplier*weight + (speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH, nil
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
// Параметры:
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("weight must be greater than 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be greater than 0")
	}
	speed := MeanSpeed(steps, duration)
	return (runningCaloriesMeanSpeedMultiplier*speed - runningCaloriesMeanSpeedShift) * weight, nil
}
