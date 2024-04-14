package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

var (
	gameColor = color.RGBA{
		R: 128,
		G: 0,
		B: 0,
		A: 255,
	}
)

type GameScene struct {
	BaseScene
}

func (s *GameScene) Update() error {
	milliseconds := time.Now().UnixMilli()
	delta := milliseconds - s.milliseconds
	s.milliseconds = milliseconds

	s.entityManager.Collect()
	s.entityManager.Update(float64(delta), s.context)

	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(gameColor)
	s.entityManager.Render(screen, s.context)
}

func NewGameScene(scene BaseScene) *GameScene {
	return &GameScene{
		BaseScene: scene,
	}
}
