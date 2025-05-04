package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	mini3d "github.com/pb82/mini3d/api"
	"log"
	"math"
	"zuzweit/api"
)

type ControlsComponent struct {
	parent  *Entity
	engine  *mini3d.Engine
	gameMap *api.Map
}

func NewControlsComponent(parent *Entity, engine *mini3d.Engine, gameMap *api.Map) *ControlsComponent {
	return parent.addComponent(&ControlsComponent{
		parent:  parent,
		engine:  engine,
		gameMap: gameMap,
	}).(*ControlsComponent)
}

func (t *ControlsComponent) Render(screen *ebiten.Image) {
	// empty
}

func (t *ControlsComponent) Update(_ float64) {

}

func (t *ControlsComponent) Type() ComponentType {
	return ControlsComponentType
}

func (t *ControlsComponent) Parent() *Entity {
	return t.parent
}

func (t *ControlsComponent) KeyInput(keys []ebiten.Key) {
	translate := t.parent.GetComponent(TranslateComponentType).(*TranslateComponent)

	for _, key := range keys {
		switch key {
		case ebiten.KeyUp:
			if api.GetCommandQueue().Empty() {
				x, y := translate.NextPosition(false)
				ahead := t.gameMap.Get(x, y)
				log.Println("position:			", math.Floor(translate.X), math.Floor(translate.Y))
				log.Println("next position:		", math.Floor(x), math.Floor(y))
				log.Println("ahead:				", ahead)

				if ahead == 1 {
					continue
				}

				api.GetCommandQueue().Push(NewAdvance(t.parent, 1, t.engine))
			}
			break
		case ebiten.KeyDown:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewAdvance(t.parent, -1, t.engine))
			}
			break
		case ebiten.KeyLeft:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewTurn(t.parent, t.engine, true))
			}
			break
		case ebiten.KeyRight:
			if api.GetCommandQueue().Empty() {
				api.GetCommandQueue().Push(NewTurn(t.parent, t.engine, false))
			}
			break
		}
	}
}
