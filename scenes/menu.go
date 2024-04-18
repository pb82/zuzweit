package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/joelschutz/stagehand"
	"image/color"
	"time"
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

func (s *MenuScene) Update() error {
	milliseconds := time.Now().UnixMilli()
	delta := milliseconds - s.milliseconds
	s.milliseconds = milliseconds

	s.entityManager.Collect()
	s.entityManager.Update(float64(delta), s.context)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		gs := NewGameScene(s.BaseScene)
		s.sm.SwitchWithTransition(gs, stagehand.NewFadeTransition[api.GameState](.05))
	}

	return nil
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(fillColor)
	s.entityManager.Render(screen, s.context)
}

func NewMenuScene(scene BaseScene) *MenuScene {
	return &MenuScene{
		BaseScene: scene,
	}
}
