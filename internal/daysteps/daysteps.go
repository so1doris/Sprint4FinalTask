package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	list := strings.Split(data, ",")
	if len(list) != 2 {
		return 0, 0, errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return 0, 0, err
	}
	if steps < 1 {
		return 0, 0, errors.New("Число шагов = 0, надо подвигаться")
	}
	t, err := time.ParseDuration(list[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, t, err := parsePackage(data)
	if err != nil {
		fmt.Errorf("Ошибка входных данных")
		return ""
	}
	if steps < 1 {
		return ""
	}
	walkingDistanceM := stepLength * float64(steps)
	walkingDistanceKm := walkingDistanceM / mInKm
	lostСalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, t)
	if err != nil {
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, walkingDistanceKm, lostСalories)
	return result
}
