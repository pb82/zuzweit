package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var (
	fillColor = color.RGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 32,
	}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(fillColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(1024, 768)
	ebiten.RunGame(game)
}
