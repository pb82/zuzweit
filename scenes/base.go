package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/joelschutz/stagehand"
	"time"
	"zuzweit/api"
	"zuzweit/ecs"
)

type BaseScene struct {
	entityManager *ecs.EntityManager
	context       *api.GameContext
	state         api.GameState
	milliseconds  int64
	delta         int64
	keys          []ebiten.Key

	sm *stagehand.SceneManager[api.GameState]
}

func (s *BaseScene) Load(state api.GameState, sm stagehand.SceneController[api.GameState]) {
	s.sm = sm.(*stagehand.SceneManager[api.GameState])
	s.state = state
}

func (s *BaseScene) Unload() api.GameState {
	return s.state
}

func (s *BaseScene) Layout(_outsideWidth, _outsideHeight int) (screenWidth, screenHeight int) {
	return api.InternalWidth, api.InternalHeight
}

func (s *BaseScene) updateCamera() {
	player := s.entityManager.GetNamedEntity("player")
	translate := player.GetComponent(ecs.TranslateComponentType).(*ecs.TranslateComponent)

	s.context.Engine.MoveCameraForward(translate.Advance)
	s.context.Engine.SetCameraPositionRelative(0, 0, 0, translate.Turn, 0)
}

func (s *BaseScene) Update() error {
	milliseconds := time.Now().UnixMilli()
	s.delta = milliseconds - s.milliseconds
	s.milliseconds = milliseconds
	s.keys = inpututil.AppendJustPressedKeys(s.keys[:0])
	s.entityManager.Collect()
	s.entityManager.Update(float64(s.delta))
	s.entityManager.KeyInput(s.keys)

	if !api.GetCommandQueue().Empty() {
		command, err := api.GetCommandQueue().Peek()
		if err != nil {
			return err
		}

		if command.Complete() {
			_, _ = api.GetCommandQueue().Pop()
		} else {
			command.Run(float64(s.delta))
		}
	}

	s.updateCamera()

	return nil
}

func NewBaseScene(em *ecs.EntityManager, context *api.GameContext) BaseScene {
	return BaseScene{
		entityManager: em,
		context:       context,
		milliseconds:  time.Now().UnixMilli(),
	}
}
