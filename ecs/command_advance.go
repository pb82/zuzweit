package ecs

import (
	mini3d "github.com/pb82/mini3d/api"
	"math"
	"zuzweit/util"
)

type Advance struct {
	player       *Entity
	increment    float64
	distance     float64
	engine       *mini3d.Engine
	speed        float64
	acceleration float64
}

func (a *Advance) Run(dt float64) {
	step := a.increment * (dt * a.speed)
	a.distance += math.Abs(step)
	a.engine.MoveCameraForward(step)
	a.speed *= a.acceleration
}

func (a *Advance) Complete() bool {
	if a.distance >= 1 {
		x, y, z, yaw, pitch := a.engine.GetCameraPosition()
		x = util.RoundFloat(x, 1)
		z = util.RoundFloat(z, 1)
		a.engine.SetCameraPositionAbsolute(x, y, z, yaw, pitch)
		return true
	}
	return false
}

func NewAdvance(player *Entity, increment float64, engine *mini3d.Engine) *Advance {
	return &Advance{
		increment:    increment,
		player:       player,
		engine:       engine,
		speed:        0.004,
		acceleration: 0.95,
	}
}
