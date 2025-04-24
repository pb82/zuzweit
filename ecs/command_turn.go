package ecs

import (
	"log"
	"math"
	api2 "zuzweit/api"
)

type Turn struct {
	player    *Entity
	increment float64
	distance  float64
}

func (a *Turn) Run(dt float64) {
	translate := a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
	step := a.increment * dt / 1000

	switch a.direction {
	case api2.North:
		translate.Y += step
	case api2.South:
		translate.Y -= step
	case api2.East:
		translate.X -= step
	case api2.West:
		translate.X += step
	}

	a.distance += math.Abs(step)
}

func (a *Turn) Complete() bool {
	if a.distance >= 1 {
		translate := a.player.GetComponent(TranslateComponentType).(*TranslateComponent)
		translate.X = math.Floor(translate.X) + .5
		translate.Y = math.Floor(translate.Y) + .5
		log.Println(translate)
		return true
	}

	return false
}

func NewTurn(player *Entity, increment float64) *Turn {
	return &Turn{
		increment: increment,
		player:    player,
		distance:  0,
	}
}
