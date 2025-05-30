package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	"image/color"
	"zuzweit/api"
	"zuzweit/ecs"
	"zuzweit/util"
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

func (s *GameScene) Load(state api.GameState, sm stagehand.SceneController[api.GameState]) {
	s.BaseScene.Load(state, sm)

	for y := 0; y < s.context.Map.H; y++ {
		for x := 0; x < s.context.Map.W; x++ {
			if s.context.Map.Get(float64(x), float64(y)) == 1 {
				cube := util.CubeWithTexture(0, 0, s.context.Atlas)
				cube.Translate(float64(x), 0, float64(y))
				s.context.Engine.AddMesh(cube)
			} else if s.context.Map.Get(float64(x), float64(y)) == 0 || s.context.Map.Get(float64(x), float64(y)) == 2 {
				cube := util.CubeWithTexture(1, 0, s.context.Atlas)
				cube.Translate(float64(x), -1, float64(y))
				s.context.Engine.AddMesh(cube)

				cube2 := util.CubeWithTexture(0, 1, s.context.Atlas)
				cube2.Translate(float64(x), 1, float64(y))
				s.context.Engine.AddMesh(cube2)
			}
		}
	}

	playerX, playerY := s.context.Map.GetPlayerStart()
	s.context.Engine.SetCameraPositionAbsolute(playerX, 0.5, playerY, 0, 0)
	t := s.entityManager.GetNamedEntity("player").GetComponent(ecs.TranslateComponentType).(*ecs.TranslateComponent)
	t.X = playerX
	t.Y = playerY
}

func (s *GameScene) Update() error {
	err := s.BaseScene.Update()
	if err != nil {
		return err
	}
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.context.Canvas.Clear()

	// first render the 3d scene
	s.context.Engine.Render(s.context.Canvas)
	screen.WritePixels(s.context.Canvas)

	// then render all other visuals
	s.entityManager.Render(screen)
}

func NewGameScene(scene BaseScene) *GameScene {
	return &GameScene{
		BaseScene: scene,
	}
}
