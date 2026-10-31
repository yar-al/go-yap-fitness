package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func checkParameters(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return errors.New("кол-во шагов должно быть положительным")
	}
	if weight <= 0 {
		return errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var paramErr = checkParameters(steps, weight, height, duration)
	if paramErr != nil {
		return 0, paramErr
	}
	var meanSpeed = MeanSpeed(steps, height, duration)
	var caloriesSpent = walkingCaloriesCoefficient * weight * meanSpeed * duration.Minutes() / minInH
	return caloriesSpent, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	var paramErr = checkParameters(steps, weight, height, duration)
	if paramErr != nil {
		return 0, paramErr
	}
	var meanSpeed = MeanSpeed(steps, height, duration)
	var caloriesSpent = weight * meanSpeed * duration.Minutes() / minInH
	return caloriesSpent, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return height * float64(steps) * stepLengthCoefficient / mInKm
}
