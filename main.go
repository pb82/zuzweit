package main

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/joelschutz/stagehand"
	mini3d "github.com/pb82/mini3d/api"
	"image/color"
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

// draw used by the mini3d engine to draw a 3d scene on a canvas
func draw(x, y int, c color.Color, userData mini3d.UserData) {
	canvas := userData.(api.Canvas)
	r, g, b, a := c.RGBA()
	pos := (y * api.InternalWidth * 4) + x*4
	canvas[pos] = byte(r >> 8)
	canvas[pos+1] = byte(g >> 8)
	canvas[pos+2] = byte(b >> 8)
	canvas[pos+3] = byte(a >> 8)
}

func main() {
	entityManager := ecs.NewEntityManager()
	engineOpts := &mini3d.EngineOptions{
		TextureAtlas: nil,
	}

	engine := mini3d.NewEngine(api.InternalWidth, api.InternalHeight, 90, draw, engineOpts)
	gameMap := api.NewDemoMap()

	player := entityManager.AddNamedEntity("player")
	ecs.NewTranslateComponent(player, 0.5, -5.5)
	ecs.NewControlsComponent(player, engine, gameMap)
	entityManager.Collect()

	context := &api.GameContext{
		Engine: engine,
		Canvas: make([]byte, api.InternalWidth*api.InternalHeight*4),
		Map:    gameMap,
	}
	state := api.GameState{}

	sceneManager := stagehand.NewSceneManager[api.GameState](
		scenes.NewGameScene(scenes.NewBaseScene(entityManager, context)),
		state)

	ebiten.SetWindowSize(api.ScreenWidth, api.ScreenHeight)
	err := ebiten.RunGame(sceneManager)
	if err != nil {
		panic(err)
	}
}
