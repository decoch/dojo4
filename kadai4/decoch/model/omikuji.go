package model

import (
	"math/rand"
	"time"
)

type FortuneSlip struct {
	Result FortuneSlipType `json:"result"`
}

type FortuneSlipType string

const (
	Daikichi FortuneSlipType = "大吉"
	Chukichi                 = "中吉"
	Kichi                    = "吉"
	Kyo                      = "凶"
)

var fortuneShips = []FortuneSlipType{
	Daikichi,
	Chukichi,
	Kichi,
	Kyo,
}

func NewFortuneShipType(t time.Time) FortuneSlipType {
	if IsFirstThreeDayOfTheNewYear(t) {
		return Daikichi
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(fortuneShips))
	return fortuneShips[i]
}

func IsFirstThreeDayOfTheNewYear(t time.Time) bool {
	return t.Month() == time.January && 1 <= t.Day() && t.Day() <= 3
}
