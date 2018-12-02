package model

import (
	"math/rand"
	"time"
)

type FortuneSlip struct {
	Result FortuneSlipType `json:"result"`
}

func DrawFortuneSlip() FortuneSlip {
	t := NewFortuneShipType(time.Now())
	return FortuneSlip{
		Result: t,
	}
}

type FortuneSlipType string

const (
	daikichi FortuneSlipType = "大吉"
	chukichi                 = "中吉"
	kichi                    = "吉"
	kyo                      = "凶"
)

var fortuneShips = []FortuneSlipType{
	daikichi,
	chukichi,
	kichi,
	kyo,
}

func NewFortuneShipType(t time.Time) FortuneSlipType {
	if IsFirstThreeDayOfTheNewYear(t) {
		return daikichi
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(fortuneShips))
	return fortuneShips[i]
}

func IsFirstThreeDayOfTheNewYear(t time.Time) bool {
	return t.Month() == time.January && 1 <= t.Day() && t.Day() <= 3
}
