package model_test

import (
	"testing"
	"time"

	"github.com/decoch/dojo4/kadai4/decoch/model"
)

func TestNewFortuneShipType(t *testing.T) {
	cases := []struct {
		date     time.Time
		shipType model.FortuneSlipType
	}{
		{date: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC), shipType: model.Daikichi},
		{date: time.Date(2018, time.January, 3, 23, 59, 59, 0, time.UTC), shipType: model.Daikichi},
	}
	for _, c := range cases {
		actual := model.NewFortuneShipType(c.date)
		if actual != c.shipType {
			t.Errorf("expected %v, actual %v.", c.shipType, actual)
		}
	}
}
