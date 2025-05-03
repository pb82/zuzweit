package ecs

import (
	mini3d "github.com/pb82/mini3d/api"
	"log"
)

type TurnLeft struct {
	player    *Entity
	engine    *mini3d.Engine
	target    float64
	increment float64
}

func (a *TurnLeft) Run(dt float64) {
	step := a.increment * (dt / 1000)
	a.engine.SetCameraPositionRelative(0, 0, 0, step, 0)
}

func (a *TurnLeft) Complete() bool {
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

func NewTurnLeft(player *Entity, engine *mini3d.Engine) *TurnLeft {
	translate := player.GetComponent(TranslateComponentType).(*TranslateComponent)

	log.Println("current: 	", translate.Compass.GetRadians())
	target := translate.Compass.TurnLeft()
	log.Println("target: 	", target)

	var increment float64

	_, _, _, yaw, _ := engine.GetCameraPosition()

	if yaw < target {
		increment = 1
		log.Println("increment: 	 +")
		log.Println("=========================")
	} else {
		log.Println("increment: 	 -")
		log.Println("=========================")
		increment = -1
	}

	turn := &TurnLeft{
		player:    player,
		engine:    engine,
		target:    target,
		increment: increment,
	}

	return turn
}
