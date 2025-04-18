package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/text/v2"
	"log"
	"zuzweit/internal"
)

type MenuComponent struct {
	parent *Entity
	root   *internal.MenuRoot
}

func NewMenuComponent(parent *Entity) *MenuComponent {
	root := &internal.MenuRoot{}
	root.PushMenu(&internal.Menu{Items: []*internal.MenuItem{
		{
			Title: "Test A",
			OnClick: func(m *internal.MenuRoot) {
				log.Println("Test A")
			},
		},
		{
			Title: "Test B",
			OnClick: func(m *internal.MenuRoot) {
				log.Println("Test B")
			},
		},
	}})

	return parent.addComponent(&MenuComponent{
		parent: parent,
		root:   root,
	}).(*MenuComponent)
}

func (t *MenuComponent) Render(screen *ebiten.Image) {
	_, err := t.root.Top()
	if err != nil {
		return
	}
}

func (t *MenuComponent) Update(_ float64) {
	// empty
}

func (t *MenuComponent) KeyInput(keys []ebiten.Key) {}

func (t *MenuComponent) Type() ComponentType {
	return MenuComponentType
}

func (t *MenuComponent) Parent() *Entity {
	return t.parent
}
