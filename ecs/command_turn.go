package ecs

import (
	mini3d "github.com/pb82/mini3d/api"
	"log"
	"math"
)

type Turn struct {
	player    *Entity
	increment float64
	target    float64
	distance  float64
	engine    *mini3d.Engine
}

func (a *Turn) Run(dt float64) {
	step := a.increment * dt / 1000
	a.distance += math.Abs(step)
	a.engine.SetCameraPositionRelative(0, 0, 0, step, 0)
}

func (a *Turn) Complete() bool {
	if a.distance >= mini3d.ToRadians(90) {
		x, y, z, yaw, pitch := a.engine.GetCameraPosition()

		degrees := yaw * 180 / math.Pi

		log.Println(degrees)

		if degrees < 0 {
			degrees = math.Ceil(degrees)
		} else {
			degrees = math.Floor(degrees)
		}

		log.Println(degrees)

		yaw = mini3d.ToRadians(degrees)

		a.engine.SetCameraPositionAbsolute(x, y, z, yaw, pitch)
		return true
	}
	return false
}

func NewTurn(player *Entity, increment float64, engine *mini3d.Engine) *Turn {
	turn := &Turn{
		increment: increment,
		player:    player,
		engine:    engine,
	}

	return turn
}
