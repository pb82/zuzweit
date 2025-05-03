package ecs

import (
	mini3d "github.com/pb82/mini3d/api"
)

type Turn struct {
	player       *Entity
	engine       *mini3d.Engine
	target       float64
	increment    float64
	speed        float64
	acceleration float64
}

func (a *Turn) Run(dt float64) {
	step := a.increment * (dt * a.speed)
	a.engine.SetCameraPositionRelative(0, 0, 0, step, 0)
	a.speed *= a.acceleration
}

func (a *Turn) Complete() bool {
	var done bool
	x, y, z, yaw, pitch := a.engine.GetCameraPosition()

	if a.increment < 0 {
		done = yaw <= a.target
	} else {
		done = yaw >= a.target
	}

	if done {
		// correct for over/undershooting
		a.engine.SetCameraPositionAbsolute(x, y, z, a.target, pitch)
	}

	return done
}

func NewTurn(player *Entity, engine *mini3d.Engine, left bool) *Turn {
	translate := player.GetComponent(TranslateComponentType).(*TranslateComponent)
	var target float64
	var increment float64

	if left {
		target = translate.Compass.TurnLeft()
	} else {
		target = translate.Compass.TurnRight()
	}

	_, _, _, yaw, _ := engine.GetCameraPosition()

	if yaw < target {
		increment = 1
	} else {
		increment = -1
	}

	turn := &Turn{
		player:       player,
		engine:       engine,
		target:       target,
		increment:    increment,
		speed:        0.01,
		acceleration: 0.9,
	}

	return turn
}
