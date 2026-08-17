package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	//splitString - слайс строк
	splitString := strings.Split(data, ",")
	if len(splitString) != 3 {
		return 0, "", 0, fmt.Errorf("lenght not 3")
	}

	//numberOfStep - количество шагов
	numberOfStep, err1 := strconv.Atoi(splitString[0])
	if err1 != nil {
		return 0, "", 0, fmt.Errorf("error: %w", err1)
	}

	//walkingDuration время прогулки
	walkingDuration, err2 := time.ParseDuration(splitString[2])
	if err2 != nil {
		return 0, "", 0, fmt.Errorf("error: %w", err2)
	}

	if walkingDuration <= 0 {
		return 0, "", 0, fmt.Errorf("time < 0")
	}

	return numberOfStep, splitString[1], walkingDuration, nil
}

func distance(steps int, height float64) float64 {
	distance := float64(steps) * height * stepLengthCoefficient
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	walkTime := duration.Hours()

	return distance / float64(walkTime)
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	return "", nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	return 0, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	return 0, nil
}
