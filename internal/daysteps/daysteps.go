package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
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
		return 0, 0, errors.New("неверный формат данных при вводе шагов")
	}
	if steps <= 0 {
		return 0, 0, errors.New("Число шагов <= 0, надо подвигаться")
	}
	t, err := time.ParseDuration(list[1])
	if err != nil {
		return 0, 0, errors.New("неверный формат данных при вводе времени")
	}
	if t <= 0 {
		return 0, 0, errors.New("неверный формат данных при вводе времени")
	}
	return steps, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, t, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps < 1 {
		fmt.Println("Число шагов <= 0, надо подвигаться")
		return ""
	}
	walkingDistanceM := stepLength * float64(steps)
	walkingDistanceKm := walkingDistanceM / mInKm
	lostСalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, t)
	if err != nil {
		log.Println(err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, walkingDistanceKm, lostСalories)
	return result
}
