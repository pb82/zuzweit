package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"zuzweit/api"
)

type ControlsComponent struct {
	parent *Entity
}

func NewControlsComponent(parent *Entity) *ControlsComponent {
	return parent.addComponent(&ControlsComponent{
		parent: parent,
	}).(*ControlsComponent)
}

func (t *ControlsComponent) Render(screen *ebiten.Image) {
	// empty
}

func (t *ControlsComponent) Update(_ float64) {

}

func (t *ControlsComponent) Type() ComponentType {
	return TranslateComponentType
}

func (t *ControlsComponent) Parent() *Entity {
	return t.parent
}

func (t *ControlsComponent) KeyInput(keys []ebiten.Key) {
	for _, key := range keys {
		switch key {
		case ebiten.KeyUp:
			advance := NewAdvance(t.parent)
			q := api.GetCommandQueue()
			q.Push(advance)
			log.Println("UP")
		}
	}
}
