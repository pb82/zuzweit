package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"image/color"
	"zuzweit/api"
)

var (
	fillColor = color.RGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 32,
	}
)

type MenuScene struct {
	BaseScene
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(fillColor)
	s.entityManager.Render(screen, s.context)
}

func (s *MenuScene) Load(state api.GameState, sm stagehand.SceneController[api.GameState]) {
	s.BaseScene.Load(state, sm)
}

func NewMenuScene(scene BaseScene) *MenuScene {
	return &MenuScene{
		BaseScene: scene,
	}
}
