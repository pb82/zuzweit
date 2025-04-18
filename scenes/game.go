package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
	mini3d "github.com/pb82/mini3d/api"
	"image/color"
	"zuzweit/api"
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

	mesh *mini3d.Mesh
}

func (s *GameScene) Load(state api.GameState, sm stagehand.SceneController[api.GameState]) {
	s.BaseScene.Load(state, sm)
	s.mesh = mini3d.ColoredCube()
	s.context.Engine.AddMesh(s.mesh)
	s.mesh.Translate(-0.5, -0.5, 0)
	s.context.Engine.SetCameraPositionRelative(0, 0, -5, mini3d.ToRadians(0), 0)
}

func (s *GameScene) Update() error {
	err := s.BaseScene.Update()
	if err != nil {
		return err
	}
	
	c := s.mesh.GetCenter()
	s.mesh.RotateXAround(1*float64(s.milliseconds)/1000, &c)
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
