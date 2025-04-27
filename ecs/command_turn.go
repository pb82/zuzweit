package ecs

import (
	"math"
)

type Turn struct {
	player    *Entity
	increment float64
	distance  float64
}

func (a *Turn) Run(dt float64) {
	_ = a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
	step := a.increment * dt / 1000
	a.distance += math.Abs(step)
}

func (a *Turn) Complete() bool {
	return false
}

func NewTurn(player *Entity, increment float64) *Turn {
	return &Turn{
		increment: increment,
		player:    player,
		distance:  0,
	}
}
