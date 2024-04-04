package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"image/color"
	"time"
	"zuzweit/ecs"

	_ "github.com/lafriks/go-tiled"
	_ "image/png"
)

//go:embed assets/tiles.png
var _tiles []byte
var tiles *ebiten.Image

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
	gameMap       *tiled.Map
	milliseconds  int64
}

func init() {
	_, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(_tiles))
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

	m, _ := tiled.LoadFile("assets/map.tmx")
	game.gameMap = m

	ebiten.SetWindowSize(1024, 768)
	ebiten.RunGame(game)
}
