package ecs

import (
	"github.com/pb82/mini3d/api"
)

type Turn struct {
	player    *Entity
	increment float64
	target    float64
}

func (a *Turn) Run(dt float64) {
	t := a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
	step := a.increment * dt / 1000
	t.Angle += step
}

func (a *Turn) Complete() bool {
	t := a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
	if a.increment > 0 {
		if t.Angle >= a.target {
			t.Angle = a.target
			return true
		}
	} else {
		if t.Angle <= a.target {
			t.Angle = a.target
			return true
		}
	}

	return false
}

func NewTurn(player *Entity, increment float64) *Turn {
	t := player.GetComponent(TranslateComponentType).(*TranslateComponent)

	turn := &Turn{
		increment: increment,
		player:    player,
		target:    t.Angle + (increment * api.ToRadians(90)),
	}

	return turn
}
