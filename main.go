package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/joelschutz/stagehand"
	"zuzweit/api"
	"zuzweit/ecs"
	"zuzweit/scenes"
)

//go:embed assets/tiles.png
var _tiles []byte
var tiles *ebiten.Image

//go:embed assets/font.ttf
var _font []byte

func init() {
	_, _, _ = ebitenutil.NewImageFromReader(bytes.NewReader(_tiles))
}

func main() {
	entityManager := ecs.NewEntityManager()
	context := &api.GameContext{}
	state := api.GameState{}

	sceneManager := stagehand.NewSceneManager[api.GameState](&scenes.MenuScene{
		BaseScene: scenes.NewBaseScene(entityManager, context),
	}, state)

	ebiten.SetWindowSize(1024, 768)
	ebiten.RunGame(sceneManager)
}
