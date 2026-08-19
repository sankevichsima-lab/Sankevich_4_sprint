package spentcalories

import (
	"fmt"
	"log"
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

	if numberOfStep <= 0 {
		return 0, "", 0, fmt.Errorf("steps must be greater than 0")
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
	distance := float64(steps) * height * stepLengthCoefficient / 1000
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
	steps, typeOfActivity, times, err1 := parseTraining(data)
	if err1 != nil {
		log.Println(err1)
		return "", fmt.Errorf("error: %w", err1)
	}

	//dist - дистанция
	dist := distance(steps, height)
	midSpeed := meanSpeed(steps, height, times)

	switch typeOfActivity {
	case "Бег":
		callories, err2 := RunningSpentCalories(steps, weight, height, times)
		if err2 != nil {
			return "", fmt.Errorf("error: %w", err2)
		}

		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfActivity,
			times.Hours(),
			dist, midSpeed,
			callories)

		return result, nil
	case "Ходьба":
		callories, err3 := WalkingSpentCalories(steps, weight, height, times)
		if err3 != nil {
			return "", fmt.Errorf("error: %w", err3)
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			typeOfActivity,
			times.Hours(),
			dist, midSpeed,
			callories)

		return result, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", typeOfActivity)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("The values  ​are incorrect.")
	}
	//midSpeed - средняя скорость
	midSpeed := meanSpeed(steps, height, duration)

	// times - время в минутах
	times := duration.Hours()

	calories := weight * midSpeed * float64(times)

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("The values  ​are incorrect.")
	}

	calories, err := RunningSpentCalories(steps, weight, height, duration)

	if err != nil {
		return 0, fmt.Errorf("error: %w", err)
	}

	result := calories * walkingCaloriesCoefficient

	return result, nil
}
