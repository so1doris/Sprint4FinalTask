package spentcalories

import (
	"errors"
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
	list := strings.Split(data, ",")
	if len(list) != 3 {
		return 0, "", 0, errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("Число шагов <= 0, надо подвигаться")
	}
	t, err := time.ParseDuration(list[2])
	if err != nil {
		return 0, "", 0, err
	}
	if t <= 0 {
		return 0, "", 0, fmt.Errorf("Время не может быть нулевым или отрецательны")
	}

	return steps, list[1], t, nil

}

func distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	return (stepLenght * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration.Seconds() <= 0 {
		return 0
	}
	actionDistance := distance(steps, height)
	return actionDistance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	/*
		Тип тренировки: Бег
		Длительность: 0.75 ч.
		Дистанция: 10.00 км.
		Скорость: 13.34 км/ч
		Сожгли калорий: 18621.75
	*/
	steps, action, t, err := parseTraining(data)
	if t.Seconds() <= 0 {
		return "", errors.New("длительность тренировки не может быть нулевой")
	}
	var spentCalories float64
	if err != nil {
		return "", err
	}
	switch action {
	case "Бег":
		spentCalories, err = RunningSpentCalories(steps, weight, height, t)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		spentCalories, err = WalkingSpentCalories(steps, weight, height, t)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		action, t.Hours(), distance(steps, height), meanSpeed(steps, height, t), spentCalories)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Неправельные параметры")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	lostCalories := (weight * averageSpeed * durationInMinutes) / minInH

	return lostCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Неправельные параметры")
	}
	durationInMinutes := duration.Minutes()
	averageSpeed := meanSpeed(steps, height, duration)
	lostCalories := (weight * averageSpeed * durationInMinutes) / minInH
	return lostCalories * walkingCaloriesCoefficient, nil
}
