package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	stepTimeStr := strings.Split(data, ",")
	if len(stepTimeStr) != 2 {
		return 0, 0, fmt.Errorf("lenght not 2")
	}

	steps, err1 := strconv.Atoi(stepTimeStr[0])

	if err1 != nil {
		return 0, 0, fmt.Errorf("error: %s", err1)
	}

	if steps < 1 {
		return 0, 0, fmt.Errorf("there cannot be less than 1 step")
	}

	walkingDuration, err2 := time.ParseDuration(stepTimeStr[1])
	if err2 != nil {
		return 0, 0, fmt.Errorf("error: %s", err2)
	}

	if walkingDuration <= 0 {
		return 0, 0, fmt.Errorf("time < 0")
	}

	return steps, walkingDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	return ""
}
