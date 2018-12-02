package model

import (
	"math/rand"
	"time"
)

type FortuneSlip struct {
	Result FortuneSlipType `json:"result"`
}

func DrawFortuneSlip(target time.Time) FortuneSlip {
	t := NewFortuneShipType(target)
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

func NewFortuneShipType(target time.Time) FortuneSlipType {
	if target.Month() == time.January && 1 <= target.Day() && target.Day() <= 3 {
		return daikichi
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(fortuneShips))
	return fortuneShips[i]
}
