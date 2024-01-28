package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
	"zuzweit/ecs"
)

var (
	fillColor = color.RGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 32,
	}
)

type Game struct {
	entityManager *ecs.EntityManager
	milliseconds  int64
}

func (g *Game) Update() error {
	milliseconds := time.Now().UnixMilli()
	delta := milliseconds - g.milliseconds
	g.milliseconds = milliseconds

	g.entityManager.Collect()
	g.entityManager.Update(float64(delta))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(fillColor)
	g.entityManager.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{
		entityManager: ecs.NewEntityManager(),
		milliseconds:  time.Now().UnixMilli(),
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.RunGame(game)
}
