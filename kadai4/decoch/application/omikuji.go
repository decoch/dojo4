package application

import (
	"time"

	"github.com/decoch/dojo4/kadai4/decoch/model"
)

type OmikujiService interface {
	Get() model.FortuneSlip
}

type omikujiService struct {
}

func NewOmikujiService() OmikujiService {
	return &omikujiService{}
}

func (o *omikujiService) Get() model.FortuneSlip {
	t := model.NewFortuneShipType(time.Now())
	return model.FortuneSlip{
		Result: t,
	}
}
